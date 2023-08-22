# GameFocus
Simple software that disables external monitors when certain game is launched.

# Build
run `build.bat` file \
*or* \
`go build -ldflags "-H=windowsgui -s -w"`

# Run
You can run this app directly. No window will show, so for now you'll need to to close the app in the Task Manger. \
The aim of this app is to automate process of turning off external displays, so you probably want to add it to the Task Scheduler

# TODO
- [x] Detect if the process is running
- [x] Disable external displays after game is launched
- [x] Enable extenral displays after game is closed
- [ ] Easily change monitored programs
- [ ] Support multiple processes
- [ ] Tray icon
- [ ] GUI for adding/removing apps
