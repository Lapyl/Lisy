## LISY

LISY is a Windows desktop app for listing, backing-up, and viewing a folder.

![alt text](https://github.com/Lapyl/Lisy/blob/main/lisy.jpg?raw=true)

## Files

- lisy.go : Go code
- Lisy.exe : Windows executable file built by using lisy.go
- go.mod and go.sum : Files generated upon running 'go mod tidy'
- lisy.jpg : Screenshot of app front-end

## Video

- YouTube video: https://youtu.be/vNfyyB-1BtI

## How to:

- Download lisy.exe from https://github.com/Lapyl/Lisy to your local folder. Alternatively, download lipy.exe.txt form https://lipy.us/docs/lipy.ee.txt to your local folder, and remove its .txt extension.
- If you want, you can add a shortcut of your local lipy.exe file.
- When needed, click lipy.exe (or its shortcut).
- For listing contents of a folder, write the folder's path in the Source folder box, and click List button. It will add two csv files of lists, in the folder.
- For backing up a folder A to another folder B, [carefully] write A's path in the Source folder box and B's path in the Backup folder box, and click Sync button.
- For a slideshow of image files from a folder, write the folder's path in the Source folder box, and click Show button. Close the app, when done. 

## To Do

- Enable Lisy to accept arguments when used through a command prompt.
- Enable Lisy to select Source and Back folders through File Explorer.
- Enable Lisy's Sync part to remember and reuse users' inputs.
- Enable Lisy's Show part to close silently.
- Extend the app for Linux, Android, and iOS.