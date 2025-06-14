# Tool4AI Terminal UI Builder

🎨 **Professional Terminal Interface Designer for AI Development**

Build beautiful, functional terminal UIs with ASCII art precision. Perfect for creating mockups, prototypes, and documentation for AI tools and applications.

![ASCII Art](https://img.shields.io/badge/ASCII-Art-brightgreen)
![Go](https://img.shields.io/badge/Go-1.19+-blue)
![Cross Platform](https://img.shields.io/badge/Platform-Cross--Platform-orange)
![License](https://img.shields.io/badge/License-MIT-green)

## ✨ Features

- 🚀 **5 Hyper Perfect Professional Templates**
  - Network Operations Center
  - Trading Floor Terminal
  - Cyber Defense Command Center
  - Data Control Station
  - Space Mission Control

- 🔧 **Core Functionality**
  - Real-time quality checking
  - Grid view with coordinates
  - Safe text replacement (never breaks structure)
  - Cross-platform single binary
  - 80-character fixed width for consistency

- 🛡️ **Safety Features**
  - Only replaces spaces - never overwrites structure
  - Auto-truncation when text is too long
  - Intelligent abbreviations for common words
  - ASCII-only character validation

## 🚀 Quick Start

### Installation

```bash
# Download the latest release
go build -o terminal-ui-builder terminal-ui-builder.go

# Or use directly
go run terminal-ui-builder.go --help
```

### Basic Usage

```bash
# Create a network operations center
./terminal-ui-builder -layout netops -h 25

# Create a trading floor interface
./terminal-ui-builder -layout trading -h 30

# Check existing ASCII art quality
./terminal-ui-builder -mode check -f examples/login_form_simple.txt

# View with debugging grid
./terminal-ui-builder -mode view -f examples/dashboard_ui.txt
```

## 🎯 Available Layouts

### Professional Templates

| Layout | Description | Use Case |
|--------|-------------|----------|
| `netops` | Network Operations Center | Infrastructure monitoring |
| `trading` | Trading Floor Terminal | Financial applications |
| `cyber` | Cyber Defense Command | Security operations |
| `data` | Data Control Station | Big data management |
| `space` | Space Mission Control | Scientific applications |

### Basic Templates

| Layout | Description | Use Case |
|--------|-------------|----------|
| `dashboard` | Simple admin dashboard | General purpose |
| `login` | Login form | Authentication |
| `box` | Custom box with title | Generic container |

## 📋 Command Reference

### Modes

- **create** - Generate new ASCII art interfaces
- **check** - Validate existing ASCII art for alignment issues
- **view** - Display with coordinate grid for debugging

### Options

```
-layout string    Layout type (default "netops")
-h int           Height of canvas (default 25)
-w int           Width for box layout (default 40, max 80)
-t string        Title for the layout
-f string        File to check or view
-mode string     Mode: create, check, view (default "create")
```

## 🌟 Examples

### Network Operations Center

```bash
./terminal-ui-builder -layout netops -h 25
```

```
+------------------------------------------------------------------------------+
| NetOps CmdCtr                                        [LIVE] [ALERT] [X]      |
+------------------------------------------------------------------------------+
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | STATUS   | | CPU      | | MEMORY   | | NETWORK  | | STORAGE  | | LATENCY  ||
| | CORE     | | 73%      | | 12GB/32GB| | 89%      | | 2.1TB/5TB| | 12ms     ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
```

### Trading Floor Terminal

```bash
./terminal-ui-builder -layout trading -h 25
```

```
+------------------------------------------------------------------------------+
| TradingFloor Terminal                           [LIVE] [RISK] [P&L] [ALERTS] |
+------------------------------------------------------------------------------+
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | PORTFOLIO| | P&L TODAY| | RISK EXPO| | MARGIN   | | EXECUTION| | MARKET   ||
| | $2.45M   | | +$21,373 | | $45,678  | | $890K    | | 1,247    | | SPY 441  ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
```

## 🔍 Quality Checking

The tool includes comprehensive quality checking:

```bash
./terminal-ui-builder -mode check -f myart.txt
```

- ✅ Vertical line alignment verification
- ⚠️ Non-ASCII character detection
- 📍 Precise error location reporting
- 🛡️ Structure integrity validation

## 📁 Project Structure

```
tool4ai_terminal-ui-builder/
├── terminal-ui-builder.go  # Main program
├── examples/              # Sample outputs
├── docs/                  # Documentation
└── releases/              # Pre-built binaries
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new layouts
5. Submit a pull request

## 📜 License

MIT License - see LICENSE file for details.

## 🎯 Use Cases

Perfect for:
- 🤖 AI tool mockups and prototypes
- 📊 System monitoring dashboards
- 💹 Financial application interfaces
- 🛡️ Security operation centers
- 🚀 Scientific mission control
- 📚 Technical documentation
- 🎓 Educational demonstrations

## 🔗 Related Projects

Part of the Tool4AI ecosystem - building better tools for AI development.

---

**Built with ❤️ for the AI development community**