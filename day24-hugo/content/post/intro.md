+++
date = "2015-09-10T02:01:35-07:00"
title = "intro"

+++

Using Hugo to create a website

Steps:
-------
- 1. brew install hugo
- 2. cd /path/to/day24-hugo
- 3. hugo new site .  (make sure the directory is empty, can't even have hidden .files)
- 4. hugo new post/intro.md
- 5. clone the theme into themes/<theme_name> (git clone https://github.com/SenjinDarashiva/hugo-uno.git themes/hugo-uno)
- 6. Change config.toml (add theme = "hugo-uno")
- 7. go to content>post>intro.md and append the blog post content and remove the draft = true line 
- 8. run the server by `hugo server -w`
- 9. Open a browser window and go to localhost:1313

