+++
title = "ExposeOrHide"
date = 2014-09-29
description = "An API for exposing or hiding properties on a model. Useful for small footprint JSON REST services"
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/ExposeOrHide"
visibility = "public"
languages = "C#"
interest_score = 2
draft = true

[extra]
post_tags = ["Old"]
+++





## About

An API for exposing or hiding properties on a model. Useful for small footprint JSON REST services

## Article

Fluent API for selectively hiding or exposing properties before serializing to JSON. Useful if you're building an API and want to return different shapes of objects to different clients without creating a million DTO classes. It's a decent pattern, but it's not particularly clever—just a filtering wrapper around serialization. Tools like GraphQL and modern API frameworks handle this kind of thing more elegantly now.

