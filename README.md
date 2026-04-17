# Pixel Counter

Browser-based tool for counting pixels by color on images. Go serves embedded static files — all logic runs client-side in JavaScript.

## Features

- Load images via file picker or drag & drop
- Pick colors directly from the image
- Adjustable threshold per color (Euclidean RGB distance)
- Manual RGB value editing
- Mask overlay with per-color and global opacity controls
- Multiple color selections (first match wins)
- Pixel counting: total, per-color, selected vs unselected
- Pixel scale: set pixel width and height separately (mm/cm/m/km)
- Area calculation in all units (mm², cm², m², km²)
- Multi-language support (English, Polish) with browser auto-detection
- Settings persisted in localStorage (language, scale, opacity)
- Zoom controls with fit-to-window

## Build

```bash
# Local build
make build

# Run
make run
# -> http://localhost:3333

# Cross-compile all platforms
make all

# Individual targets
make windows   # dist/pixelcounter-windows-amd64.exe
make linux     # dist/pixelcounter-linux-amd64
make mac       # dist/pixelcounter-darwin-amd64 + arm64
```

## Requirements

- Go 1.22+

## License

CC BY-NC 4.0 — free for non-commercial use. See [LICENSE](LICENSE) for details.
