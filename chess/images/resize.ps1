# Get a list of all PNG files in the current directory
$files = Get-ChildItem -Filter *.png

# Set the desired width and height for resizing
$width = 100
$height = 100

# Resize each image
foreach ($file in $files) {
    $image = [System.Drawing.Image]::FromFile($file.FullName)
    $newImage = $image.GetThumbnailImage($width, $height, $null, [System.IntPtr]::Zero)
    $newImage.Save($file.FullName)
}

Write-Host "Images resized successfully."
