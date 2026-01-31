+++
title = "String.Latinize"
date = 1351777024
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/String.Latinize"
visibility = "public"
languages = "C#"
interest_score = 0
+++

## About

Make any unicode string conform to [a..zA..Z], keep whitespace and symbols/numbers
## Original README

String.Latinize
===============

Make any unicode string conform to [a..zA..Z], keep whitespace and symbols/numbers

Install using Package Manager Console: 

    PM> Install-Package Latinize

Example:
	
    //using Latinize;
	Console.WriteLine("äæåâáàãÄÆÅÂÁÀÃ çÇ éèêëeÉÈÊËE øöØÖ".Latinize());
	//=>               aaaaaaaAAAAAAA cC eeeeeEEEEE ooOO
