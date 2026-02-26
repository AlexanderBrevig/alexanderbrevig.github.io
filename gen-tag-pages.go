package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Post struct {
	Title       string
	Date        string
	Description string
	Permalink   string
}

func main() {
	// Create tags directory if it doesn't exist
	tagsDir := "content/tags"
	os.MkdirAll(tagsDir, 0755)

	blogDir := "content/blog"
	files, err := os.ReadDir(blogDir)
	if err != nil {
		log.Fatal(err)
	}

	tagMap := make(map[string][]Post)
	tagCounts := make(map[string]int)

	// Parse all blog posts
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") && file.Name() != "_index.md" {
			filePath := filepath.Join(blogDir, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			text := string(content)
			parts := strings.SplitN(text, "+++", 3)
			if len(parts) < 2 {
				continue
			}

			fm := parts[1]
			title := extractField(fm, "title")
			date := extractField(fm, "date")
			desc := extractField(fm, "description")
			tags := extractTags(fm)

			if title == "" || len(tags) == 0 {
				continue
			}

			postName := strings.TrimSuffix(file.Name(), ".md")
			post := Post{
				Title:       title,
				Date:        date,
				Description: desc,
				Permalink:   "/blog/" + postName + "/",
			}

			for _, tag := range tags {
				tagMap[tag] = append(tagMap[tag], post)
				tagCounts[tag]++
			}
		}
	}

	// Generate tag listing page (_index.md for /tags/)
	generateTagsListingPage(tagCounts)

	// Generate individual tag pages
	for tag, posts := range tagMap {
		// Sort by date descending
		sort.Slice(posts, func(i, j int) bool {
			return posts[i].Date > posts[j].Date
		})
		generateTagPage(tag, posts)
	}

	fmt.Printf("Generated %d tag pages\n", len(tagMap))
}

func generateTagsListingPage(tagCounts map[string]int) {
	// Sort tags by name
	var tags []string
	for tag := range tagCounts {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	var content strings.Builder
	content.WriteString(`+++
title = "Tags"
template = "tags-list.html"
render = true

[extra]
+++

Click on a tag to view all articles in that category.

`)

	// Create a listing
	for _, tag := range tags {
		count := tagCounts[tag]
		countStr := "1 article"
		if count > 1 {
			countStr = fmt.Sprintf("%d articles", count)
		}
		slug := slugify(tag)
		content.WriteString(fmt.Sprintf("- [%s](/tags/%s/) (%s)\n", tag, slug, countStr))
	}

	err := os.WriteFile("content/tags/_index.md", []byte(content.String()), 0644)
	if err != nil {
		log.Printf("Error writing tags index: %v\n", err)
	}
}

func generateTagPage(tag string, posts []Post) {
	tagSlug := slugify(tag)
	tagDir := filepath.Join("content/tags", tagSlug)
	os.MkdirAll(tagDir, 0755)

	var content strings.Builder
	content.WriteString(fmt.Sprintf(`+++
title = "%s"
template = "tag-page.html"
render = true

[extra]
tag_name = "%s"
+++

`, tag, tag))

	// List all posts with this tag (will add tag links via Tera template)
	for _, post := range posts {
		content.WriteString(fmt.Sprintf("- [%s](%s) — %s\n", post.Title, post.Permalink, post.Date))
		if post.Description != "" {
			content.WriteString(fmt.Sprintf("  > %s\n", post.Description))
		}
	}

	filePath := filepath.Join(tagDir, "_index.md")
	err := os.WriteFile(filePath, []byte(content.String()), 0644)
	if err != nil {
		log.Printf("Error writing tag page for %s: %v\n", tag, err)
	}
}

func extractField(fm string, field string) string {
	lines := strings.Split(fm, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, field) {
			idx := strings.Index(line, "=")
			if idx != -1 {
				value := strings.TrimSpace(line[idx+1:])
				value = strings.Trim(value, `"`)
				return value
			}
		}
	}
	return ""
}

func extractTags(fm string) []string {
	lines := strings.Split(fm, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "tags") {
			idx := strings.Index(line, "=")
			if idx != -1 {
				value := strings.TrimSpace(line[idx+1:])
				if !strings.HasPrefix(value, "[") {
					continue
				}
				value = strings.Trim(value, "[]")
				var tags []string
				for _, t := range strings.Split(value, ",") {
					t = strings.TrimSpace(t)
					t = strings.Trim(t, `"`)
					if t != "" {
						tags = append(tags, t)
					}
				}
				return tags
			}
		}
	}
	return []string{}
}

func slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	// Replace spaces and special chars with hyphens
	var result strings.Builder
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			result.WriteRune(ch)
		} else if ch == ' ' || ch == '_' || ch == '-' || ch == '/' || ch == '(' || ch == ')' {
			// Only add hyphen if last char isn't already a hyphen
			if result.Len() > 0 {
				last := result.String()
				if !strings.HasSuffix(last, "-") {
					result.WriteRune('-')
				}
			}
		}
		// Skip other special characters
	}
	// Remove trailing hyphens
	s = strings.TrimSuffix(result.String(), "-")
	return s
}
