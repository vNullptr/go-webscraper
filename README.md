# WebScraper 📦

A lightweight, modular, and extensible Go-based web scraping library.
Designed with simplicity in mind — built for developers who want full control over requests, parsing, and extraction… without the bloat.

## 🚀 Overview

WebScraper is an ongoing project aimed at becoming a flexible scraping toolkit written in Go.
The goal is to provide:

- Clean abstractions
- A plug-and-play parser system
- A minimal learning curve
- High performance using Go’s concurrency features

## ✨ Current Features

- URL Fetching — simple, reliable download layer.
- Basic HTML Parsing — uses Go’s html parser under the hood.
- Element Extraction Helpers — early utilities for scraping text, attributes, and nodes.

## 🔧 Technical Focus (Work in Progress)

- Cleaning and normalizing the golang.org/x/net/html node tree.
- Designing my own custom internal DOM representation for easier traversal and querying.
- Reworking the structure to be intuitive, documented, and stable.

## 🧭 Roadmap / Planned Features
### Networking

- Custom headers, cookies, user agents
- Request timeout configuration ( currently testing ) 
- Retry logic & error strategy patterns ( currently testing )

### Parsing & Extraction

- Pluggable parser modules (HTML, XML, JSON)
- Querying system for selecting nodes easily
- Better utilities for attributes, inner text, traversal

### Performance

- Native concurrency enabled scraping

## ⚠️ The current repository structure is provisional, a full cleanup and reorganization is planned.
