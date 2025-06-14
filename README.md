# Tool4AI Terminal UI Builder

🤖 **AI-First Terminal Interface Designer**

**Designed specifically for AI assistants like Claude Code, ChatGPT, and other AI development tools.**

This tool enables AI assistants to generate professional, pixel-perfect terminal interfaces instantly. Instead of struggling with manual ASCII art alignment, AI can now create complex dashboards, forms, and UIs with a single command.

## 🎯 Why AI-First?

- **Perfect for AI assistants** - Simple commands generate complex interfaces
- **No manual alignment** - AI doesn't need to worry about spacing or structure
- **Consistent output** - Every generation is perfectly formatted
- **Instant prototyping** - AI can create UI mockups in seconds
- **Professional quality** - Production-ready terminal interfaces

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

### For AI Assistants

**Claude Code, ChatGPT, and other AI tools can use this directly:**

```bash
# AI can generate professional dashboards instantly
./terminal-ui-builder -layout netops -h 25

# Create trading interfaces for financial apps
./terminal-ui-builder -layout trading -h 30

# Generate login forms
./terminal-ui-builder -layout login -h 18

# AI can validate generated ASCII art
./terminal-ui-builder -mode check -f user_interface.txt
```

### For Human Developers

```bash
# Download the latest release
go build -o terminal-ui-builder terminal-ui-builder.go

# Or use directly
go run terminal-ui-builder.go --help
```

### AI-Friendly Commands

AI assistants can use these simple patterns:

```bash
# Generate specific layouts
./terminal-ui-builder -layout [netops|trading|cyber|data|space] -h [height]

# Quality check any ASCII art
./terminal-ui-builder -mode check -f [filename]

# Debug with grid overlay
./terminal-ui-builder -mode view -f [filename]
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

**Perfect for AI Assistants:**
- 🤖 **Instant UI generation** - AI creates complex interfaces with single commands
- 📝 **Documentation assistance** - AI generates visual mockups for specifications
- 🎓 **Teaching tool** - AI can demonstrate UI concepts with real examples
- 🔧 **Rapid prototyping** - AI builds UI mockups during development discussions
- 📊 **System design** - AI creates architecture diagrams and dashboards

**Perfect for Human Developers:**
- 📊 System monitoring dashboards
- 💹 Financial application interfaces
- 🛡️ Security operation centers
- 🚀 Scientific mission control
- 📚 Technical documentation

## 🔗 Related Projects

Part of the Tool4AI ecosystem - building better tools for AI development.

---

## 🤖 AI Integration Examples

**Claude Code usage:**
```bash
# AI can instantly create a network operations dashboard
./terminal-ui-builder -layout netops -h 25 > network_dashboard.txt

# AI can validate and fix ASCII art alignment issues
./terminal-ui-builder -mode check -f my_interface.txt
```

**ChatGPT integration:**
```bash
# Generate professional trading interfaces
./terminal-ui-builder -layout trading -h 30

# Create space mission control interfaces
./terminal-ui-builder -layout space -h 25
```

**Perfect for AI workflows:**
- No manual ASCII art struggles
- Consistent, professional output
- Instant validation and debugging
- Cross-platform compatibility

---

**Built with ❤️ for AI assistants and the AI development community**