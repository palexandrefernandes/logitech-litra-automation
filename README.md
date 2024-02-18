# Logitech Litra camera light automation

Automating litra lightbar to automatically turn on when a camera is being used in linux.

## Usage
At the moment the script is only tracking one camera and litra device. You can
change which camera the script is targeting by updating the camera device path
(following V4L2 framework standard) in the `main.go` file.

If you want to run this script on you will need to manually do so.