+++
title = "InstanceListPattern"
date = 2012-10-14
description = "Easily make your object maintain an internal list of instances, relies on the programmer to Dispose();"
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/InstanceListPattern"
visibility = "public"
languages = "C#"
interest_score = 2
draft = true

[extra]
post_tags = ["Old"]
+++





## About

Easily make your object maintain an internal list of instances, relies on the programmer to Dispose();

## Article

This is a clever little pattern for automatically keeping track of every instance of a class. Inherit from `InstanceList<T>` and boom—you've got a static list of all living instances. It's the kind of thing that's useful for specific scenarios (game objects, entity tracking, etc.) but really only if you're diligent about calling Dispose(). Without that, you'll leak memory hard. It's a nice idea, but it's solving a niche problem and it's easy to shoot yourself in the foot with it.

