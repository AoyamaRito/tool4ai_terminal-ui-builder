package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Canvas with fixed 80 character width
const CANVAS_WIDTH = 80

type AsciiStudio struct {
	height int
	lines  [][]rune
}

func NewAsciiStudio(height int) *AsciiStudio {
	lines := make([][]rune, height)
	for i := range lines {
		lines[i] = make([]rune, CANVAS_WIDTH)
		for j := range lines[i] {
			lines[i][j] = ' '
		}
	}
	return &AsciiStudio{height: height, lines: lines}
}

// Safe text addition - only replaces spaces
func (as *AsciiStudio) AddTextSafe(x, y int, text string) {
	if y < 0 || y >= as.height {
		return
	}
	
	availableWidth := 0
	for i := x; i < CANVAS_WIDTH; i++ {
		if as.lines[y][i] == ' ' {
			availableWidth++
		} else {
			break
		}
	}
	
	truncatedText := truncateText(text, availableWidth)
	runes := []rune(truncatedText)
	for i, r := range runes {
		if x+i >= 0 && x+i < CANVAS_WIDTH && as.lines[y][x+i] == ' ' {
			as.lines[y][x+i] = r
		}
	}
}

func truncateText(text string, maxWidth int) string {
	if len(text) <= maxWidth {
		return text
	}
	if maxWidth <= 3 {
		return text[:maxWidth]
	}
	return text[:maxWidth-3] + "..."
}

func abbreviateText(text string) string {
	abbreviations := map[string]string{
		"Network Operations": "NetOps",
		"Command Center": "CmdCtr",
		"Security": "Sec",
		"Analytics": "Stats",
		"Management": "Mgmt",
		"Configuration": "Config",
		"Administration": "Admin",
		"Monitoring": "Monitor",
		"Dashboard": "Dash",
		"Trading Floor": "TradeFl",
		"Cyber Defense": "CyberDef",
		"Data Control": "DataCtrl",
		"Mission Control": "MsnCtrl",
		"Performance": "Perf",
		"Intelligence": "Intel",
		"Emergency": "Emerg",
		"Response": "Resp",
	}
	
	for full, abbrev := range abbreviations {
		text = strings.ReplaceAll(text, full, abbrev)
	}
	return text
}

// Drawing functions
func (as *AsciiStudio) DrawFrame() {
	for i := 0; i < CANVAS_WIDTH; i++ {
		as.lines[0][i] = '+'
		as.lines[as.height-1][i] = '+'
		if i > 0 && i < CANVAS_WIDTH-1 {
			as.lines[0][i] = '-'
			as.lines[as.height-1][i] = '-'
		}
	}
	for row := 1; row < as.height-1; row++ {
		as.lines[row][0] = '|'
		as.lines[row][CANVAS_WIDTH-1] = '|'
	}
}

func (as *AsciiStudio) DrawBox(x, y, w, h int) {
	if x < 0 || y < 0 || x+w > CANVAS_WIDTH || y+h > as.height {
		return
	}
	
	as.lines[y][x] = '+'
	as.lines[y][x+w-1] = '+'
	as.lines[y+h-1][x] = '+'
	as.lines[y+h-1][x+w-1] = '+'
	
	for i := x + 1; i < x+w-1; i++ {
		as.lines[y][i] = '-'
		as.lines[y+h-1][i] = '-'
	}
	
	for i := y + 1; i < y+h-1; i++ {
		as.lines[i][x] = '|'
		as.lines[i][x+w-1] = '|'
	}
}

func (as *AsciiStudio) DrawHLine(x, y, length int) {
	if y < 0 || y >= as.height {
		return
	}
	for i := 0; i < length && x+i < CANVAS_WIDTH; i++ {
		if x+i >= 0 {
			as.lines[y][x+i] = '-'
		}
	}
}

func (as *AsciiStudio) DrawFullHLine(y int) {
	if y < 0 || y >= as.height {
		return
	}
	as.lines[y][0] = '+'
	for i := 1; i < CANVAS_WIDTH-1; i++ {
		as.lines[y][i] = '-'
	}
	as.lines[y][CANVAS_WIDTH-1] = '+'
}

func (as *AsciiStudio) DrawVLine(x, y, length int) {
	if x < 0 || x >= CANVAS_WIDTH {
		return
	}
	for i := 0; i < length && y+i < as.height; i++ {
		if y+i >= 0 {
			as.lines[y+i][x] = '|'
		}
	}
}

func (as *AsciiStudio) String() string {
	var result strings.Builder
	for _, line := range as.lines {
		result.WriteString(string(line))
		result.WriteRune('\n')
	}
	return result.String()
}

// Hyper Perfect Templates
func createNetworkOps(height int) *AsciiStudio {
	as := NewAsciiStudio(height)
	as.DrawFrame()
	
	as.AddTextSafe(2, 1, abbreviateText("NetOps Command Center"))
	as.AddTextSafe(CANVAS_WIDTH-25, 1, "[LIVE] [ALERT] [X]")
	as.DrawFullHLine(2)
	
	// Status boxes
	boxWidth := 12
	for i := 0; i < 6; i++ {
		x := 2 + i*(boxWidth+1)
		as.DrawBox(x, 4, boxWidth, 4)
		switch i {
		case 0:
			as.AddTextSafe(x+2, 5, "STATUS")
			as.AddTextSafe(x+2, 6, "CORE")
		case 1:
			as.AddTextSafe(x+2, 5, "CPU")
			as.AddTextSafe(x+2, 6, "73%")
		case 2:
			as.AddTextSafe(x+2, 5, "MEMORY")
			as.AddTextSafe(x+2, 6, "12GB/32GB")
		case 3:
			as.AddTextSafe(x+2, 5, "NETWORK")
			as.AddTextSafe(x+2, 6, "89%")
		case 4:
			as.AddTextSafe(x+2, 5, "STORAGE")
			as.AddTextSafe(x+2, 6, "2.1TB/5TB")
		case 5:
			as.AddTextSafe(x+2, 5, "LATENCY")
			as.AddTextSafe(x+2, 6, "12ms")
		}
	}
	
	// Main panels
	as.DrawBox(2, 9, 30, 8)
	as.AddTextSafe(4, 10, "LIVE TRAFFIC FLOW")
	as.AddTextSafe(4, 11, "US-WEST [====] 892 Mbps")
	as.AddTextSafe(4, 12, "US-EAST [===] 634 Mbps")
	as.AddTextSafe(4, 13, "EU-WEST [==] 423 Mbps")
	
	as.DrawBox(34, 9, 44, 8)
	as.AddTextSafe(36, 10, abbreviateText("INCIDENT RESPONSE"))
	as.AddTextSafe(36, 11, "HIGH: SSL Cert Expiring 3d")
	as.AddTextSafe(36, 12, "MED: High mem usage DB-02")
	as.AddTextSafe(36, 13, "LOW: CDN cache miss >15%")
	
	as.AddTextSafe(2, height-2, "Ops: 4 engineers | Global: 99.97% uptime")
	
	return as
}

func createTradingFloor(height int) *AsciiStudio {
	as := NewAsciiStudio(height)
	as.DrawFrame()
	
	as.AddTextSafe(2, 1, abbreviateText("TradingFloor Terminal"))
	as.AddTextSafe(CANVAS_WIDTH-30, 1, "[LIVE] [RISK] [P&L] [ALERTS]")
	as.DrawFullHLine(2)
	
	// Stats boxes
	boxWidth := 12
	for i := 0; i < 6; i++ {
		x := 2 + i*(boxWidth+1)
		as.DrawBox(x, 4, boxWidth, 4)
		switch i {
		case 0:
			as.AddTextSafe(x+2, 5, "PORTFOLIO")
			as.AddTextSafe(x+2, 6, "$2.45M")
		case 1:
			as.AddTextSafe(x+2, 5, "P&L TODAY")
			as.AddTextSafe(x+2, 6, "+$21,373")
		case 2:
			as.AddTextSafe(x+2, 5, "RISK EXPO")
			as.AddTextSafe(x+2, 6, "$45,678")
		case 3:
			as.AddTextSafe(x+2, 5, "MARGIN")
			as.AddTextSafe(x+2, 6, "$890K")
		case 4:
			as.AddTextSafe(x+2, 5, "EXECUTION")
			as.AddTextSafe(x+2, 6, "1,247")
		case 5:
			as.AddTextSafe(x+2, 5, "MARKET")
			as.AddTextSafe(x+2, 6, "SPY 441.25")
		}
	}
	
	// Trading panels
	as.DrawBox(2, 9, 25, 8)
	as.AddTextSafe(4, 10, "ACTIVE POSITIONS")
	as.AddTextSafe(4, 11, "SPY  |100 |+1250|2.1%")
	as.AddTextSafe(4, 12, "AAPL |200 |+890 |1.8%")
	as.AddTextSafe(4, 13, "TSLA |-50 |-230 |0.9%")
	
	as.DrawBox(29, 9, 25, 8)
	as.AddTextSafe(31, 10, "PENDING ORDERS")
	as.AddTextSafe(31, 11, "LMT|AAPL|50|185.50")
	as.AddTextSafe(31, 12, "STP|TSLA|25|245.00")
	as.AddTextSafe(31, 13, "MKT|SPY |75|MKT")
	
	as.DrawBox(56, 9, 22, 8)
	as.AddTextSafe(58, 10, "MARKET DEPTH - SPY")
	as.AddTextSafe(58, 11, "2,500|441.22|1,800")
	as.AddTextSafe(58, 12, "3,200|441.23|2,100")
	as.AddTextSafe(58, 13, "1,900|441.24|1,500")
	
	as.AddTextSafe(2, height-2, "Trader: J.Smith | Connectivity: LOW LATENCY")
	
	return as
}

func createCyberCommand(height int) *AsciiStudio {
	as := NewAsciiStudio(height)
	as.DrawFrame()
	
	as.AddTextSafe(2, 1, abbreviateText("CyberDefense Command Center"))
	as.AddTextSafe(CANVAS_WIDTH-25, 1, "[DEFCON 3] [ACTIVE]")
	as.DrawFullHLine(2)
	
	// Security status boxes
	boxWidth := 12
	for i := 0; i < 6; i++ {
		x := 2 + i*(boxWidth+1)
		as.DrawBox(x, 4, boxWidth, 4)
		switch i {
		case 0:
			as.AddTextSafe(x+2, 5, "SECURITY")
			as.AddTextSafe(x+2, 6, "DEFCON: 3")
		case 1:
			as.AddTextSafe(x+2, 5, "THREATS")
			as.AddTextSafe(x+2, 6, "Crit: 0")
		case 2:
			as.AddTextSafe(x+2, 5, "INCIDENTS")
			as.AddTextSafe(x+2, 6, "P1: 0 P2: 2")
		case 3:
			as.AddTextSafe(x+2, 5, "NETWORK")
			as.AddTextSafe(x+2, 6, "2.4 Gbps")
		case 4:
			as.AddTextSafe(x+2, 5, "ENDPOINTS")
			as.AddTextSafe(x+2, 6, "12,456")
		case 5:
			as.AddTextSafe(x+2, 5, "INTEL")
			as.AddTextSafe(x+2, 6, "47 feeds")
		}
	}
	
	// Threat panels
	as.DrawBox(2, 9, 35, 8)
	as.AddTextSafe(4, 10, "LIVE THREAT MAP")
	as.AddTextSafe(4, 11, "US-EAST [***] 23 attacks")
	as.AddTextSafe(4, 12, "US-WEST [**] 12 attacks")
	as.AddTextSafe(4, 13, "EU-WEST [*] 8 attacks")
	as.AddTextSafe(4, 14, "ASIA-PAC [****] 34 attacks")
	
	as.DrawBox(39, 9, 39, 8)
	as.AddTextSafe(41, 10, abbreviateText("INCIDENT TIMELINE"))
	as.AddTextSafe(41, 11, "12:47 [P2] Suspicious PowerShell")
	as.AddTextSafe(41, 12, "12:45 [P3] Failed login x50")
	as.AddTextSafe(41, 13, "12:42 [P2] DLL injection")
	as.AddTextSafe(41, 14, "12:38 [P4] Port scan detected")
	
	as.AddTextSafe(2, height-2, "SOC Analysts: 8 active | MTTR: 45 min")
	
	return as
}

func createDataControl(height int) *AsciiStudio {
	as := NewAsciiStudio(height)
	as.DrawFrame()
	
	as.AddTextSafe(2, 1, abbreviateText("DataMaster Control Station"))
	as.AddTextSafe(CANVAS_WIDTH-25, 1, "[SYNC] [BACKUP] [LIVE]")
	as.DrawFullHLine(2)
	
	// Data pipeline boxes
	boxWidth := 12
	for i := 0; i < 6; i++ {
		x := 2 + i*(boxWidth+1)
		as.DrawBox(x, 4, boxWidth, 4)
		switch i {
		case 0:
			as.AddTextSafe(x+2, 5, "SOURCES")
			as.AddTextSafe(x+2, 6, "3.24M/s")
		case 1:
			as.AddTextSafe(x+2, 5, "PIPELINE")
			as.AddTextSafe(x+2, 6, "98.7% OK")
		case 2:
			as.AddTextSafe(x+2, 5, "CLUSTERS")
			as.AddTextSafe(x+2, 6, "Active")
		case 3:
			as.AddTextSafe(x+2, 5, "QUERIES")
			as.AddTextSafe(x+2, 6, "Run: 23")
		case 4:
			as.AddTextSafe(x+2, 5, "STREAMS")
			as.AddTextSafe(x+2, 6, "4.2M/s")
		case 5:
			as.AddTextSafe(x+2, 5, "ML JOBS")
			as.AddTextSafe(x+2, 6, "Training")
		}
	}
	
	// Data panels
	as.DrawBox(2, 9, 25, 8)
	as.AddTextSafe(4, 10, "DATA QUALITY")
	as.AddTextSafe(4, 11, "[PASS] Schema validation")
	as.AddTextSafe(4, 12, "[WARN] 2 null fields")
	as.AddTextSafe(4, 13, "[PASS] 0.001% dupes")
	as.AddTextSafe(4, 14, "[ALERT] 3 outliers")
	
	as.DrawBox(29, 9, 25, 8)
	as.AddTextSafe(31, 10, "REPLICATION")
	as.AddTextSafe(31, 11, "Master->Slave: 0.2s")
	as.AddTextSafe(31, 12, "Backup->Archive: 1.2s")
	as.AddTextSafe(31, 13, "Cross-Region: 45ms")
	as.AddTextSafe(31, 14, "Health: GREEN")
	
	as.DrawBox(56, 9, 22, 8)
	as.AddTextSafe(58, 10, "ML PIPELINE")
	as.AddTextSafe(58, 11, "Features: 2.4M")
	as.AddTextSafe(58, 12, "Training: XGBoost")
	as.AddTextSafe(58, 13, "Serving: 12K req/s")
	as.AddTextSafe(58, 14, "Accuracy: 95%")
	
	as.AddTextSafe(2, height-2, "Data Engineers: 6 | Processing: 3.24M/sec")
	
	return as
}

func createSpaceMission(height int) *AsciiStudio {
	as := NewAsciiStudio(height)
	as.DrawFrame()
	
	as.AddTextSafe(2, 1, abbreviateText("Mission Control Center - Mars Rover"))
	as.AddTextSafe(CANVAS_WIDTH-20, 1, "[T+2847d] [ACTIVE]")
	as.DrawFullHLine(2)
	
	// Mission status boxes
	boxWidth := 12
	for i := 0; i < 6; i++ {
		x := 2 + i*(boxWidth+1)
		as.DrawBox(x, 4, boxWidth, 4)
		switch i {
		case 0:
			as.AddTextSafe(x+2, 5, "POSITION")
			as.AddTextSafe(x+2, 6, "-14.57,137")
		case 1:
			as.AddTextSafe(x+2, 5, "POWER SYS")
			as.AddTextSafe(x+2, 6, "847W")
		case 2:
			as.AddTextSafe(x+2, 5, "SCIENCE")
			as.AddTextSafe(x+2, 6, "1,247 samp")
		case 3:
			as.AddTextSafe(x+2, 5, "COMMS")
			as.AddTextSafe(x+2, 6, "2.4 kbps")
		case 4:
			as.AddTextSafe(x+2, 5, "HEALTH")
			as.AddTextSafe(x+2, 6, "45.2C")
		case 5:
			as.AddTextSafe(x+2, 5, "WEATHER")
			as.AddTextSafe(x+2, 6, "-73C/12C")
		}
	}
	
	// Mission panels
	as.DrawBox(2, 9, 30, 8)
	as.AddTextSafe(4, 10, "MISSION TIMELINE")
	as.AddTextSafe(4, 11, "Sol 2847: Science ops")
	as.AddTextSafe(4, 12, "09:15 Drive waypoint")
	as.AddTextSafe(4, 13, "10:30 Deploy drill")
	as.AddTextSafe(4, 14, "11:45 Collect sample")
	as.AddTextSafe(4, 15, "13:20 Analyze SAM")
	
	as.DrawBox(34, 9, 22, 8)
	as.AddTextSafe(36, 10, "INSTRUMENT STATUS")
	as.AddTextSafe(36, 11, "MAHLI: [READY]")
	as.AddTextSafe(36, 12, "MastCam: [ACTIVE]")
	as.AddTextSafe(36, 13, "ChemCam: [READY]")
	as.AddTextSafe(36, 14, "SAM: [PROCESSING]")
	as.AddTextSafe(36, 15, "RAD: [CONTINUOUS]")
	
	as.DrawBox(58, 9, 20, 8)
	as.AddTextSafe(60, 10, "DISCOVERIES")
	as.AddTextSafe(60, 11, "Organic: 23 det")
	as.AddTextSafe(60, 12, "Rock age: 3.7By")
	as.AddTextSafe(60, 13, "pH: 6.8 neutral")
	as.AddTextSafe(60, 14, "Minerals: 47")
	as.AddTextSafe(60, 15, "Tubes: 18/38")
	
	as.AddTextSafe(2, height-2, "Sol 2847 | Signal Delay: 14m 23s")
	
	return as
}

// Checking functions (same as before)
func (as *AsciiStudio) CheckAlignment() []string {
	var issues []string
	lines := strings.Split(strings.TrimSpace(as.String()), "\n")
	
	verticalPositions := make(map[int][]int)
	for i, line := range lines {
		for j, char := range line {
			if char == '|' {
				verticalPositions[j] = append(verticalPositions[j], i)
			}
		}
	}
	
	for col, rows := range verticalPositions {
		if len(rows) > 1 {
			for i := 0; i < len(rows)-1; i++ {
				gap := rows[i+1] - rows[i]
				if gap > 1 {
					missingRows := []int{}
					for row := rows[i] + 1; row < rows[i+1]; row++ {
						if col < len(lines[row]) {
							char := rune(lines[row][col])
							allowedChars := "-+|"
							if col > 2 && col < len(lines[row])-2 {
								allowedChars += " "
							}
							if !strings.ContainsRune(allowedChars, char) {
								missingRows = append(missingRows, row+1)
							}
						}
					}
					if len(missingRows) > 0 {
						issues = append(issues, fmt.Sprintf("Column %d: Vertical alignment gap at rows %v", col, missingRows))
					}
				}
			}
		}
	}
	
	for i, line := range lines {
		for j, char := range line {
			if char > 127 {
				issues = append(issues, fmt.Sprintf("Line %d, Col %d: Non-ASCII character '%c' found", i+1, j+1, char))
			}
		}
	}
	
	return issues
}

func (as *AsciiStudio) ShowGrid() {
	lines := strings.Split(strings.TrimSpace(as.String()), "\n")
	
	fmt.Println("Grid View with Coordinates:")
	fmt.Println(strings.Repeat("=", CANVAS_WIDTH))
	
	fmt.Print("   ")
	for i := 0; i < CANVAS_WIDTH; i++ {
		if i%10 == 0 {
			fmt.Print(i / 10)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	
	fmt.Print("   ")
	for i := 0; i < CANVAS_WIDTH; i++ {
		fmt.Print(i % 10)
	}
	fmt.Println()
	
	for i, line := range lines {
		fmt.Printf("%2d %s\n", i, line)
	}
}

func main() {
	var (
		mode     = flag.String("mode", "create", "Mode: create, check, view")
		layout   = flag.String("layout", "netops", "Layout: netops, trading, cyber, data, space, dashboard, login, box")
		height   = flag.Int("h", 25, "Height of the canvas")
		width    = flag.Int("w", 40, "Width for simple box (max 80)")
		title    = flag.String("t", "", "Title for the layout")
		filename = flag.String("f", "", "File to check or view")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ASCII Studio COMPLETE - All Templates in One Tool\n\n")
		fmt.Fprintf(os.Stderr, "⚠️  HYPER PERFECT SERIES INCLUDED:\n")
		fmt.Fprintf(os.Stderr, "    • Network Operations Center (netops)\n")
		fmt.Fprintf(os.Stderr, "    • Trading Floor Terminal (trading)\n")
		fmt.Fprintf(os.Stderr, "    • Cyber Defense Command (cyber)\n")
		fmt.Fprintf(os.Stderr, "    • Data Control Station (data)\n")
		fmt.Fprintf(os.Stderr, "    • Space Mission Control (space)\n")
		fmt.Fprintf(os.Stderr, "    • Plus: dashboard, login, box\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nModes:\n")
		fmt.Fprintf(os.Stderr, "  create    - Create new ASCII art\n")
		fmt.Fprintf(os.Stderr, "  check     - Check existing ASCII art for issues\n")
		fmt.Fprintf(os.Stderr, "  view      - View with grid overlay\n")
		fmt.Fprintf(os.Stderr, "\nLayouts (for create mode):\n")
		fmt.Fprintf(os.Stderr, "  netops    - Network Operations Center\n")
		fmt.Fprintf(os.Stderr, "  trading   - Trading Floor Terminal\n")
		fmt.Fprintf(os.Stderr, "  cyber     - Cyber Defense Command\n")
		fmt.Fprintf(os.Stderr, "  data      - Data Control Station\n")
		fmt.Fprintf(os.Stderr, "  space     - Space Mission Control\n")
		fmt.Fprintf(os.Stderr, "  dashboard - Simple admin dashboard\n")
		fmt.Fprintf(os.Stderr, "  login     - Login form\n")
		fmt.Fprintf(os.Stderr, "  box       - Custom box (use with -w and -t)\n")
		fmt.Fprintf(os.Stderr, "\nCanvas is always 80 characters wide\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -layout netops -h 30\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -layout trading -h 35\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -layout cyber -h 25\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -mode check -f myart.txt\n", os.Args[0])
	}
	
	flag.Parse()
	
	switch *mode {
	case "create":
		var as *AsciiStudio
		
		switch *layout {
		case "netops":
			as = createNetworkOps(*height)
		case "trading":
			as = createTradingFloor(*height)
		case "cyber":
			as = createCyberCommand(*height)
		case "data":
			as = createDataControl(*height)
		case "space":
			as = createSpaceMission(*height)
		case "dashboard":
			// Simple dashboard from previous version
			as = NewAsciiStudio(*height)
			as.DrawFrame()
			as.AddTextSafe(2, 1, "DASH")
			as.AddTextSafe(CANVAS_WIDTH-15, 1, "User: admin")
			as.DrawFullHLine(2)
			sidebarWidth := 20
			as.DrawVLine(sidebarWidth, 2, *height-3)
			as.AddTextSafe(2, 4, "NAV")
		case "login":
			// Simple login from previous version
			as = NewAsciiStudio(*height)
			as.DrawFrame()
			as.AddTextSafe((CANVAS_WIDTH-10)/2, 1, "LOGIN FORM")
			as.DrawFullHLine(2)
		case "box":
			// Simple box
			as = NewAsciiStudio(*height)
			startX := (CANVAS_WIDTH - *width) / 2
			if startX < 0 {
				startX = 0
				*width = CANVAS_WIDTH
			}
			as.DrawBox(startX, 0, *width, *height)
			if *title != "" {
				titleX := startX + (*width-len(*title))/2
				as.AddTextSafe(titleX, 1, *title)
			}
		default:
			fmt.Fprintf(os.Stderr, "Unknown layout: %s\n", *layout)
			os.Exit(1)
		}
		
		fmt.Print(as.String())
		
	case "check":
		var content string
		
		if *filename != "" {
			data, err := os.ReadFile(*filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(1)
			}
			content = string(data)
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			content = strings.Join(lines, "\n")
		}
		
		contentLines := strings.Split(content, "\n")
		height := len(contentLines)
		as := NewAsciiStudio(height)
		
		for i, line := range contentLines {
			if i < height {
				runes := []rune(line)
				for j, r := range runes {
					if j < CANVAS_WIDTH {
						as.lines[i][j] = r
					}
				}
			}
		}
		
		issues := as.CheckAlignment()
		fmt.Printf("ASCII Art Check Results:\n")
		fmt.Printf("========================\n")
		
		if len(issues) == 0 {
			fmt.Println("✅ No issues found! ASCII art looks good.")
		} else {
			fmt.Printf("❌ Found %d issues:\n\n", len(issues))
			for i, issue := range issues {
				fmt.Printf("%d. %s\n", i+1, issue)
			}
		}
		
	case "view":
		var content string
		
		if *filename != "" {
			data, err := os.ReadFile(*filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(1)
			}
			content = string(data)
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			content = strings.Join(lines, "\n")
		}
		
		contentLines := strings.Split(content, "\n")
		height := len(contentLines)
		as := NewAsciiStudio(height)
		
		for i, line := range contentLines {
			if i < height {
				runes := []rune(line)
				for j, r := range runes {
					if j < CANVAS_WIDTH {
						as.lines[i][j] = r
					}
				}
			}
		}
		
		as.ShowGrid()
		
	default:
		fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", *mode)
		fmt.Fprintf(os.Stderr, "Available modes: create, check, view\n")
		os.Exit(1)
	}
}