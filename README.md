# Tool4AI Terminal UI Builder

🤖 **AI-First Terminal Interface Designer**

**📋 Enable Claude Code to create pixel-perfect ASCII art UI designs instantly!**

**Designed specifically for AI assistants like Claude Code, ChatGPT, and other AI development tools.**

**Perfect workflow: AI creates → Humans view → AI codes**

1. 🤖 **AI easily creates ASCII art UIs** - Claude Code generates complex interfaces with perfect alignment
2. 👁️ **Humans can view without special viewers** - Plain text format readable anywhere 
3. 💻 **AI can code from these designs** - Claude Code uses the ASCII designs as blueprints for implementation

By using this tool, Claude Code can generate professional, accurate ASCII art terminal interfaces without alignment issues. Instead of struggling with manual ASCII art spacing problems, Claude Code can now create complex dashboards, forms, and UIs with perfect precision using a single command.

## 🎯 Complete Design-to-Code Workflow

**For AI Assistants:**
- **🤖 Easy ASCII UI creation** - Generate complex interfaces without alignment struggles
- **📐 Perfect ASCII art alignment** - Pixel-perfect terminal UIs every time
- **💻 Blueprint for coding** - Use generated designs as implementation guides
- **⚡ Instant generation** - Simple commands create professional interfaces

**For Human Developers:**
- **👁️ Universal viewing** - No special viewers needed, readable in any text editor
- **📋 Clear design documents** - Perfect for specifications and documentation
- **🤝 AI-Human collaboration** - AI creates designs, humans review, AI implements
- **✅ Consistent communication** - Same format for design and development phases

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

#### macOS
```bash
# Intel Mac
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-darwin-amd64 -o terminal-ui-builder
chmod +x terminal-ui-builder

# Apple Silicon Mac
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-darwin-arm64 -o terminal-ui-builder
chmod +x terminal-ui-builder
```

#### Linux
```bash
# AMD64
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-linux-amd64 -o terminal-ui-builder
chmod +x terminal-ui-builder

# ARM64
curl -L https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-linux-arm64 -o terminal-ui-builder
chmod +x terminal-ui-builder
```

#### Windows
```powershell
# PowerShell - AMD64
Invoke-WebRequest https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-windows-amd64.exe -OutFile terminal-ui-builder.exe

# PowerShell - ARM64
Invoke-WebRequest https://github.com/AoyamaRito/tool4ai_terminal-ui-builder/raw/main/releases/terminal-ui-builder-windows-arm64.exe -OutFile terminal-ui-builder.exe
```

#### Build from Source
```bash
git clone https://github.com/AoyamaRito/tool4ai_terminal-ui-builder.git
cd tool4ai_terminal-ui-builder
go build -o terminal-ui-builder terminal-ui-builder.go
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

## 🌟 Live Demos

### 🚀 Complex Network Operations Center

AI can generate sophisticated monitoring dashboards instantly:

```bash
./terminal-ui-builder -layout netops -h 35
```

```
+------------------------------------------------------------------------------+
| NetOps CmdCtr                                        [LIVE] [ALERT] [X]      |
+------------------------------------------------------------------------------+
|                                                                              |
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | STATUS   | | CPU      | | MEMORY   | | NETWORK  | | STORAGE  | | LATENCY  ||
| | CORE     | | 73%      | | 12GB/32GB| | 89%      | | 2.1TB/5TB| | 12ms     ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
|                                                                              |
| +----------------------------+  +------------------------------------------+ |
| | LIVE TRAFFIC FLOW          |  | INCIDENT RESPONSE                        | |
| | US-WEST [====] 892 Mbps    |  | HIGH: SSL Cert Expiring 3d               | |
| | US-EAST [===] 634 Mbps     |  | MED: High mem usage DB-02                | |
| | EU-WEST [==] 423 Mbps      |  | LOW: CDN cache miss >15%                 | |
| +----------------------------+  +------------------------------------------+ |
| Ops: 4 engineers | Global: 99.97% uptime                                     |
+------------------------------------------------------------------------------+
```

### 💹 Professional Trading Terminal

Perfect for financial AI applications:

```bash
./terminal-ui-builder -layout trading -h 30
```

```
+------------------------------------------------------------------------------+
| TradingFloor Terminal                           [LIVE] [RISK] [P&L] [ALERTS] |
+------------------------------------------------------------------------------+
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | PORTFOLIO| | P&L TODAY| | RISK EXPO| | MARGIN   | | EXECUTION| | MARKET   ||
| | $2.45M   | | +$21,373 | | $45,678  | | $890K    | | 1,247    | | SPY 441  ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| +-----------------------+  +-----------------------+  +--------------------+ |
| | ACTIVE POSITIONS      |  | PENDING ORDERS        |  | MARKET DEPTH - SPY | |
| | SPY  |100 |+1250|2.1% |  | LMT|AAPL|50|185.50    |  | 2,500|441.22|1,800 | |
| | AAPL |200 |+890 |1.8% |  | STP|TSLA|25|245.00    |  | 3,200|441.23|2,100 | |
| +-----------------------+  +-----------------------+  +--------------------+ |
+------------------------------------------------------------------------------+
```

### 🛡️ Cyber Security Command Center

AI-generated security operations dashboard:

```bash
./terminal-ui-builder -layout cyber -h 25
```

```
+------------------------------------------------------------------------------+
| CyberDefense CmdCtr                                  [DEFCON 3] [ACTIVE]     |
+------------------------------------------------------------------------------+
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | SECURITY | | THREATS  | | INCIDENTS| | NETWORK  | | ENDPOINTS| | INTEL    ||
| | DEFCON: 3| | Crit: 0  | | P1: 0    | | 2.4 Gbps | | 12,456   | | 47 feeds ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| +---------------------------------+  +-------------------------------------+ |
| | LIVE THREAT MAP                 |  | INCIDENT TIMELINE                   | |
| | US-EAST [***] 23 attacks        |  | 12:47 [P2] Suspicious PowerShell    | |
| | ASIA-PAC [****] 34 attacks      |  | 12:45 [P3] Failed login x50         | |
| +---------------------------------+  +-------------------------------------+ |
+------------------------------------------------------------------------------+
```

### 🚀 Space Mission Control

Science & research AI applications:

```bash
./terminal-ui-builder -layout space -h 25
```

```
+------------------------------------------------------------------------------+
| MsnCtrl Center - Mars Rover                               [T+2847d] [ACTIVE] |
+------------------------------------------------------------------------------+
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| | POSITION | | POWER SYS| | SCIENCE  | | COMMS    | | HEALTH   | | WEATHER  ||
| | -14.57...|  | 847W     | | 1,247 ...| | 2.4 kbps | | 45.2C    | | -73C/12C ||
| +----------+ +----------+ +----------+ +----------+ +----------+ +----------+|
| +----------------------------+  +--------------------+  +------------------+ |
| | MISSION TIMELINE           |  | INSTRUMENT STATUS  |  | DISCOVERIES      | |
| | Sol 2847: Science ops      |  | MAHLI: [READY]     |  | Organic: 23 det  | |
| | 09:15 Drive waypoint       |  | MastCam: [ACTIVE]  |  | Tubes: 18/38     | |
| +----------------------------+  +--------------------+  +------------------+ |
+------------------------------------------------------------------------------+
```

### ⚡ One-Command Generation

**AI assistants can create any of these complex interfaces with a single command:**

```bash
# Network operations dashboard
./terminal-ui-builder -layout netops -h 35 > network_dashboard.txt

# Trading floor terminal  
./terminal-ui-builder -layout trading -h 30 > trading_terminal.txt

# Security operations center
./terminal-ui-builder -layout cyber -h 25 > security_center.txt
```

**Perfect for:**
- 🤖 AI documentation generation
- 📋 System design discussions  
- 🎓 Teaching interface concepts
- 🔧 Rapid prototyping sessions

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