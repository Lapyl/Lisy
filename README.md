## LISY

LISY is a Windows desktop app for listing, backing-up, and viewing a folder.

![alt text](https://github.com/Lapyl/Lisy/blob/main/lisy.jpg?raw=true)

## Files

- lisy.go : Go code
- Lisy.exe.txt : Windows executable file built by using lisy.go (with .txt added for security)
- go.mod and go.sum : Files generated upon running 'go mod tidy'
- lisy.jpg : Screenshot of app front-end

## Video

- YouTube video: https://youtu.be/4ySbtfIrCfg

## How to:

- Download lisy.exe.txt from https://github.com/Lapyl/Lisy or https://lipy.us/docs to your local folder.
- Remove .txt from the end of the downloaded file's name.
- If you want, you may add a shortcut of your local lisy.exe file.
- When needed, click lisy.exe (or its shortcut).
- For listing contents of a folder, fetch or write the folder's path in the Source folder box, and click List button. It will add two csv files of lists, in the folder.
- For backing up a folder A to another folder B, [carefully] fetch or write A's path in the Source folder box and fetch or write B's path in the Backup folder box, and click Sync button.
- For a slideshow of image files from a folder, fetch or write the folder's path in the Source folder box, and click Show button. Close the app, when done.

## To Do

- Enable Lisy to accept arguments when used through a command prompt.
- Enable Lisy's Sync part to remember and reuse users' inputs.
- Enable Lisy's Show part to close silently.
- Extend the app for Linux, Android, and iOS.