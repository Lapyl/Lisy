## Project

Lisy stands for 'List Insp Sync Ying'. It is a Windows desktop app, developed by Lalit A Patel and Laxmi L Patel, and submitted to Digital Jam Hackathon 2025.

## Files on this Repository

- lisy.go : Go code
- Lisy.exe : Windows executable file built by using lisy.go

## Links

- Project Submission on DevPost: https://devpost.com/software/lisy-windows-desktop-app-for-files-management
- Project Video on YouTube: https://youtu.be/oRSAzVVCHJ0
- GitHub repository: https://github.com/Lapyl/Lisy
- Hackathon Information on Devpost: https://digital-jam.devpost.com/

## Description on DevPost

- Inspiration
- Keeping an inventory of digital files on computers is of paramount importance. We often like to have a quick preview and inspection of pictures stored in a folder. All of us routinely need to create back up replicas of drives or folders.
- What it does
- Lisy’s List module allows us to prepare lists of contents in a selected folder. Lisy’s Insp module allows us to slide through files in a folder. Lisy’s Sync module can be used for to creating back up replicas of drives or folders. Lisy’ Ying is for a little fun.
- How we built it
- Lisy has been written using Google’s Go programming language, which is efficient and easy to use. After reviewing pros and cons of various GUI approaches, we have used Gioui package and some Html approach for providing a user-friendly graphical user interface. 
- Challenges we ran into
- We found that Gioui package is not so good as Fyne package for canvases with images.
- Accomplishments that we're proud of
- Windows executable file Lisy.exe that we have built can be immensely useful in our and others' routine tasks.  
- What we learned
- We have sharpened our skills in Go programming languages. We now understand different GUI approaches available for Go.
- What's next for Lisy
- We will enhance the code to: Allow Lisy List and Sync modules to remember and reuse users' inputs.  Allow Insp module to show pictures on an application window rather than as a wallpaper and to change the speed of sliding.

## Narration in video

- Hi, I am Lalit Patel.
- And, I am Laxmi Patel.
- We have developed a Windows desktop app named Lisy.
- Keeping an inventory of digital files on computers is of paramount importance.
- Lisy’s List module allows us to prepare lists of contents in a selected folder.
- We often like to have a quick preview and inspection of pictures stored in a folder.
- Lisy’s Insp module allows us to slide through files in a folder.
- All of us routinely need to create back up replicas of drives or folders. 
- Lisy’s Sync module can be used for that purpose.
- How about some fun?
- Lisy’ Ying can help.
- Lisy has been written using Google’s Go programming language.
- Go is efficient, and easy to use.
- Here is the main Go code used for Lisy.
- ‘Go Mod Tidy’ command brings in necessary libraries.
- ‘Go Run Dot’ command runs the code while developing.
- ‘Go Build’ compiles and builds a Windows executable file.
- Lisy’s application window is simple.
- Top part has 4 buttons for the 4 modules.
- Below that is a text box for writing inputs.
- Then there is a progress bar and an image block. 
- For the List part, you specify your folder’s path.
- Clicking List quickly prepares two csv files, for subfolders and files.
- Each list has item name, file size, and modified date.
- For the Insp part too, you specify a folder’s path.
- Clicking Insp presents pictures one by one. 
- Future versions of Lisy will allow us to configure how pictures are presented.
- The Sync part needs paths of source and target folders.  
- Clicking Sync backs up the source folder to the target folder. 
- Sync uses Robocopy and is pretty fast.
- The Ying part needs you to specify seconds for heating the rock picture.
- Clicking Ying changes size and color of the image.
- It can be a little fun to watch this. 
- Thank you.
- Thank you.