# Usage Guide

## Installation

### Build from Source

```bash
git clone https://github.com/AoyamaRito/tool4ai_terminal-ui-builder.git
cd tool4ai_terminal-ui-builder
go build -o terminal-ui-builder terminal-ui-builder.go
```

### Cross-Platform Build

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o releases/terminal-ui-builder-linux-amd64 terminal-ui-builder.go

# macOS AMD64
GOOS=darwin GOARCH=amd64 go build -o releases/terminal-ui-builder-darwin-amd64 terminal-ui-builder.go

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o releases/terminal-ui-builder-darwin-arm64 terminal-ui-builder.go

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o releases/terminal-ui-builder-windows-amd64.exe terminal-ui-builder.go
```

## Command Line Interface

### Basic Syntax

```bash
terminal-ui-builder [mode] [options]
```

### Modes

#### Create Mode (Default)

Generate new ASCII art interfaces:

```bash
# Basic usage
terminal-ui-builder -layout netops

# With custom height
terminal-ui-builder -layout trading -h 30

# Custom box with title
terminal-ui-builder -layout box -w 50 -h 15 -t "My Application"
```

#### Check Mode

Validate existing ASCII art:

```bash
# Check a file
terminal-ui-builder -mode check -f examples/dashboard_ui.txt

# Check from stdin
cat myart.txt | terminal-ui-builder -mode check
```

#### View Mode

Display with coordinate grid:

```bash
# View with grid overlay
terminal-ui-builder -mode view -f examples/login_form_simple.txt
```

## Options Reference

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `-mode` | string | `create` | Mode: create, check, view |
| `-layout` | string | `netops` | Layout type |
| `-h` | int | `25` | Canvas height |
| `-w` | int | `40` | Box width (max 80) |
| `-t` | string | `""` | Title text |
| `-f` | string | `""` | Input file path |

## Examples

### Professional Dashboards

```bash
# Network Operations Center
terminal-ui-builder -layout netops -h 25 > netops-dashboard.txt

# Trading Floor
terminal-ui-builder -layout trading -h 30 > trading-terminal.txt

# Cyber Security
terminal-ui-builder -layout cyber -h 25 > security-center.txt
```

### Quality Assurance

```bash
# Check for issues
terminal-ui-builder -mode check -f netops-dashboard.txt

# Debug with grid
terminal-ui-builder -mode view -f netops-dashboard.txt
```

### Pipeline Integration

```bash
# Generate and check in one pipeline
terminal-ui-builder -layout data -h 25 | \
terminal-ui-builder -mode check
```

## Tips and Tricks

### 1. Optimal Heights by Layout

- **netops**: 20-25 rows
- **trading**: 25-35 rows  
- **cyber**: 20-30 rows
- **data**: 25-35 rows
- **space**: 25-30 rows

### 2. Text Truncation

The tool automatically truncates long text:
- Preserves meaning with "..." suffix
- Respects word boundaries when possible
- Uses intelligent abbreviations

### 3. Safe Editing

- Text only replaces spaces
- Never overwrites structural characters
- Maintains perfect alignment

### 4. Cross-Platform Testing

```bash
# Test on different systems
./terminal-ui-builder -layout netops > test.txt
cat test.txt  # Verify display
```

## Troubleshooting

### Common Issues

1. **Text appears cut off**
   - Solution: Increase canvas height with `-h`

2. **Alignment looks wrong**
   - Solution: Use `-mode check` to identify issues
   - Ensure monospace font in terminal

3. **Special characters display incorrectly**
   - Solution: Use ASCII-only mode
   - Check terminal encoding (UTF-8 recommended)

### Getting Help

```bash
# Show all options
terminal-ui-builder --help

# Check version info
terminal-ui-builder -version  # (if implemented)
```