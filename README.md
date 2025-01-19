# About

fork of diary Project, which I had written as a practice side project on 30 Dec 2024. 

This is a CLI diary, to allow me to journal (directly from the command-line) what I have done at the end of a workday.

The program provides prompts: What did you accomplish today? What went well? What didn't go well? What could be improved? 

The program accepts the command-line input, and outputs a timestamped entry in a Markdown file. 
![](CLI%20Diary%20Screenshot1.png)
![](CLI%20Diary%20Screenshot2.png)
<br>

# Additional - diaryProject
_30 Dec 2024 0936hrs_
I have created a diaryProject: this is a command-line interface diary program, written in Go, which allows the user to create, update, read or delete a diary entry file, directly from the command-line

The purpose of this CLI app is to scratch my own itch: I often focus on my tasks, and end the day without any reflection of how the day went. So the day ends without any retrospection, which is a wasted learning opportunity.

Having this CLI app serves to force me to reflect on what went well, what didn't go so well, and what could be improved for the next day. 

I will dogfood this app for a while, and make iterative tweaks and improvements to the project along the way. 

# Changelog
- _30Dec24 1257hrs_: refactored codebase: main now very clean, while helper functions all put into a separate file.
- _30Dec24 1651hrs_: further refactored codebase, with createEntry and updateEntry helper functions further refactored into smaller, more modular helpers. Also improved the UX by adding some basic formatting e.g. displaying the Entry after the entry has been successfully created/updated. added a gitignore, so I don\'t accidentally upload my diary entries on Github! 


# TODOs
- [x] To refactor codebase: a lot of repetition 
