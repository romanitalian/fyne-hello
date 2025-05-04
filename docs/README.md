# Documentation

## Version History

### v0.0.1
- Initial release
- Basic button and label functionality
- Simple window management
- Material Design-inspired UI

## Development

### Building
```bash
go build
```

### Running
```bash
go run main.go
```

### Packaging
For macOS:
```bash
fyne package -os darwin -icon Icon.png
```

For Windows:
```bash
fyne package -os windows -icon Icon.png
```

For Linux:
```bash
fyne package -os linux -icon Icon.png
```

## Architecture

The application follows a simple structure:
- `main.go`: Entry point and UI setup
- `FyneApp.toml`: Application metadata
- `Icon.png`: Application icon

## Dependencies

- fyne.io/fyne/v2 v2.6.0
- Go 1.24 or later 