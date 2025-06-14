// ASCII UI Instructor - ASCII UIの作成指示を生成するツール
// 
// 使い方:
//   go run ascii-ui-instructor.go [options]
//
// インストール:
//   go build -o ascii-ui-instructor ascii-ui-instructor.go
//   sudo mv ascii-ui-instructor /usr/local/bin/
//
// 使用例:
//   ascii-ui-instructor -layout dashboard
//   ascii-ui-instructor -layout login
//   ascii-ui-instructor -mode instructions -layout dashboard
//
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type UIElement struct {
	Type        string
	Name        string
	Position    Position
	Size        Size
	Content     []string
	Children    []UIElement
	Style       string
	Priority    int
	Actions     []UIAction
	Inputs      []UIInput
}

type UIAction struct {
	Type        string
	Trigger     string
	Description string
	Handler     string
}

type UIInput struct {
	Type        string
	Name        string
	Validation  string
	Required    bool
	Placeholder string
}

type Position struct {
	X, Y string
}

type Size struct {
	Width, Height string
}

type Layout struct {
	Name        string
	Width       int
	Height      int
	Elements    []UIElement
	Description string
}

type InstructionGenerator struct {
	layout       *Layout
	instructions []string
}

func NewInstructionGenerator(layout *Layout) *InstructionGenerator {
	return &InstructionGenerator{
		layout:       layout,
		instructions: []string{},
	}
}

func (ig *InstructionGenerator) Generate() []string {
	ig.instructions = []string{}
	
	ig.addInstruction(fmt.Sprintf("# ASCII UI Instructions for: %s", ig.layout.Name))
	ig.addInstruction(fmt.Sprintf("Canvas Size: %d x %d characters", ig.layout.Width, ig.layout.Height))
	ig.addInstruction("")
	
	if ig.layout.Description != "" {
		ig.addInstruction(fmt.Sprintf("Description: %s", ig.layout.Description))
		ig.addInstruction("")
	}
	
	ig.addInstruction("## Step-by-Step Instructions:")
	ig.addInstruction("")
	
	ig.addInstruction("### 1. Create the main frame")
	ig.addInstruction(fmt.Sprintf("- Draw a border box using +, -, and | characters"))
	ig.addInstruction(fmt.Sprintf("- Frame dimensions: %d x %d", ig.layout.Width, ig.layout.Height))
	ig.addInstruction("")
	
	step := 2
	for i, element := range ig.layout.Elements {
		ig.generateElementInstructions(element, step, i+1)
		step++
	}
	
	ig.addInstruction("### Final Steps:")
	ig.addInstruction("- Review all alignments")
	ig.addInstruction("- Ensure all vertical lines connect properly")
	ig.addInstruction("- Check that text doesn't overflow boundaries")
	ig.addInstruction("- Verify consistent spacing and padding")
	
	return ig.instructions
}

func (ig *InstructionGenerator) generateElementInstructions(element UIElement, step int, index int) {
	ig.addInstruction(fmt.Sprintf("### %d. Add %s: %s", step, element.Type, element.Name))
	
	// Add actions and inputs information
	if len(element.Actions) > 0 {
		ig.addInstruction("**Expected Actions:**")
		for _, action := range element.Actions {
			ig.addInstruction(fmt.Sprintf("- %s: %s (trigger: %s)", action.Type, action.Description, action.Trigger))
			if action.Handler != "" {
				ig.addInstruction(fmt.Sprintf("  Handler: %s", action.Handler))
			}
		}
	}
	
	if len(element.Inputs) > 0 {
		ig.addInstruction("**Expected Inputs:**")
		for _, input := range element.Inputs {
			required := ""
			if input.Required {
				required = " (required)"
			}
			ig.addInstruction(fmt.Sprintf("- %s: %s%s", input.Name, input.Type, required))
			if input.Validation != "" {
				ig.addInstruction(fmt.Sprintf("  Validation: %s", input.Validation))
			}
			if input.Placeholder != "" {
				ig.addInstruction(fmt.Sprintf("  Placeholder: %s", input.Placeholder))
			}
		}
	}
	
	switch element.Type {
	case "header":
		ig.addInstruction(fmt.Sprintf("- Position: %s from left, row %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Add text: \"%s\"", strings.Join(element.Content, " ")))
		if element.Style != "" {
			ig.addInstruction(fmt.Sprintf("- Style: %s", element.Style))
		}
		
	case "box":
		ig.addInstruction(fmt.Sprintf("- Position: %s, %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Size: %s x %s", element.Size.Width, element.Size.Height))
		ig.addInstruction("- Draw using +, -, and | characters")
		if len(element.Content) > 0 {
			ig.addInstruction("- Content:")
			for _, content := range element.Content {
				ig.addInstruction(fmt.Sprintf("  - %s", content))
			}
		}
		
	case "separator":
		ig.addInstruction(fmt.Sprintf("- Draw horizontal line at row %s", element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Width: %s", element.Size.Width))
		ig.addInstruction("- Use - characters, connect to frame with +")
		
	case "panel":
		ig.addInstruction(fmt.Sprintf("- Create panel at position %s, %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Dimensions: %s x %s", element.Size.Width, element.Size.Height))
		if element.Name != "" {
			ig.addInstruction(fmt.Sprintf("- Title: \"%s\"", element.Name))
		}
		if len(element.Children) > 0 {
			ig.addInstruction("- Panel contents:")
			for _, child := range element.Children {
				ig.addInstruction(fmt.Sprintf("  - %s: %s", child.Type, strings.Join(child.Content, ", ")))
			}
		}
		
	case "statusbar":
		ig.addInstruction(fmt.Sprintf("- Position: bottom row - %s", element.Position.Y))
		ig.addInstruction("- Content sections:")
		for i, content := range element.Content {
			alignment := "left"
			if i == len(element.Content)-1 {
				alignment = "right"
			} else if i > 0 {
				alignment = "center"
			}
			ig.addInstruction(fmt.Sprintf("  - %s aligned: \"%s\"", alignment, content))
		}
		
	case "menu":
		ig.addInstruction(fmt.Sprintf("- Position: %s aligned, row %s", element.Style, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Menu items: %s", strings.Join(element.Content, " ")))
		
	case "progressbar":
		ig.addInstruction(fmt.Sprintf("- Position: %s, %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Width: %s characters", element.Size.Width))
		ig.addInstruction("- Use [ ] for boundaries, = or █ for filled portions")
		if len(element.Content) > 0 {
			ig.addInstruction(fmt.Sprintf("- Label: %s", element.Content[0]))
		}
		
	case "table":
		ig.addInstruction(fmt.Sprintf("- Position: %s, %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Size: %s x %s", element.Size.Width, element.Size.Height))
		ig.addInstruction("- Structure:")
		ig.addInstruction("  - Use | for column separators")
		ig.addInstruction("  - Use - for row separators")
		ig.addInstruction("  - Align column data consistently")
		if len(element.Content) > 0 {
			ig.addInstruction(fmt.Sprintf("  - Headers: %s", strings.Join(element.Content, " | ")))
		}
		
	case "button":
		ig.addInstruction(fmt.Sprintf("- Position: %s, %s", element.Position.X, element.Position.Y))
		ig.addInstruction(fmt.Sprintf("- Format: [%s]", strings.Join(element.Content, " ")))
		
	default:
		ig.addInstruction(fmt.Sprintf("- Type: %s", element.Type))
		ig.addInstruction(fmt.Sprintf("- Position: %s, %s", element.Position.X, element.Position.Y))
		if element.Size.Width != "" || element.Size.Height != "" {
			ig.addInstruction(fmt.Sprintf("- Size: %s x %s", element.Size.Width, element.Size.Height))
		}
		if len(element.Content) > 0 {
			ig.addInstruction(fmt.Sprintf("- Content: %s", strings.Join(element.Content, ", ")))
		}
	}
	
	ig.addInstruction("")
}

func (ig *InstructionGenerator) addInstruction(instruction string) {
	ig.instructions = append(ig.instructions, instruction)
}

func generateSilentOutput(layout *Layout) []string {
	var output []string
	
	// First, generate the ASCII art representation
	output = append(output, "以下のASCII UIを作成してください:")
	output = append(output, "")
	output = append(output, fmt.Sprintf("画面名: %s", layout.Name))
	output = append(output, fmt.Sprintf("サイズ: %d x %d 文字", layout.Width, layout.Height))
	output = append(output, "")
	output = append(output, "要求仕様:")
	output = append(output, "1. 枠線は +, -, | を使用して描画")
	output = append(output, "2. すべての要素は指定された位置に配置")
	output = append(output, "3. テキストは枠内に収まるように配置")
	output = append(output, "")
	output = append(output, "UI要素:")
	
	// List all UI elements
	for i, element := range layout.Elements {
		output = append(output, fmt.Sprintf("%d. %s (%s)", i+1, element.Name, element.Type))
		output = append(output, fmt.Sprintf("   位置: %s, %s", element.Position.X, element.Position.Y))
		if element.Size.Width != "" || element.Size.Height != "" {
			output = append(output, fmt.Sprintf("   サイズ: %s x %s", element.Size.Width, element.Size.Height))
		}
		if len(element.Content) > 0 {
			output = append(output, fmt.Sprintf("   内容: %s", strings.Join(element.Content, ", ")))
		}
		output = append(output, "")
	}
	
	// Add interactive elements section
	output = append(output, "ASCII UIの下に以下の形式でINTERACTIVE_ELEMENTSセクションを追加:")
	output = append(output, "")
	output = append(output, "INTERACTIVE_ELEMENTS:")
	
	for _, element := range layout.Elements {
		if len(element.Actions) > 0 {
			for j, action := range element.Actions {
				buttonName := ""
				if len(element.Content) > j && strings.Contains(element.Content[j], "[") {
					buttonName = element.Content[j]
				} else if len(element.Content) > 0 {
					buttonName = element.Content[0]
				} else {
					buttonName = element.Name
				}
				
				transition := action.Handler
				if transition == "" {
					transition = "(画面遷移なし)"
				} else if strings.Contains(transition, "close") || strings.Contains(transition, "exit") {
					transition = "(アプリケーション終了)"
				} else {
					transition = strings.ReplaceAll(transition, "()", ".txt")
					transition = strings.ReplaceAll(transition, "open", "")
					transition = strings.ReplaceAll(transition, "show", "")
				}
				
				output = append(output, fmt.Sprintf("%s - %s: %s → %s", buttonName, action.Trigger, action.Description, transition))
			}
		}
		
		if len(element.Inputs) > 0 {
			for _, input := range element.Inputs {
				output = append(output, fmt.Sprintf("%s - input: %s入力 → (画面遷移なし)", input.Name, input.Type))
			}
		}
	}
	
	return output
}

func analyzeASCIIArt(content string) *Layout {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	height := len(lines)
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	
	layout := &Layout{
		Name:   "Analyzed Layout",
		Width:  width,
		Height: height,
	}
	
	elements := []UIElement{}
	
	if height > 2 && strings.Contains(lines[0], "+") && strings.Contains(lines[0], "-") {
		title := extractTextFromLine(lines[1])
		if title != "" {
			elements = append(elements, UIElement{
				Type:     "header",
				Name:     "Main Title",
				Position: Position{X: "2", Y: "1"},
				Content:  []string{title},
			})
		}
	}
	
	for i, line := range lines {
		if i > 0 && i < height-1 && strings.Count(line, "+") >= 2 && strings.Contains(line, "-") {
			elements = append(elements, UIElement{
				Type:     "separator",
				Name:     fmt.Sprintf("Separator at row %d", i),
				Position: Position{Y: fmt.Sprintf("%d", i)},
				Size:     Size{Width: "full"},
			})
		}
	}
	
	layout.Elements = elements
	return layout
}

func extractTextFromLine(line string) string {
	line = strings.Trim(line, "| ")
	parts := strings.Split(line, "|")
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}
	return strings.TrimSpace(line)
}

func createSampleLayouts() map[string]*Layout {
	layouts := make(map[string]*Layout)
	
	layouts["dashboard"] = &Layout{
		Name:        "System Dashboard",
		Width:       80,
		Height:      25,
		Description: "A comprehensive system monitoring dashboard",
		Elements: []UIElement{
			{
				Type:     "header",
				Name:     "Title Bar",
				Position: Position{X: "2", Y: "1"},
				Content:  []string{"System Monitor v2.0"},
				Style:    "left-aligned",
			},
			{
				Type:     "header",
				Name:     "Status Indicators",
				Position: Position{X: "right-15", Y: "1"},
				Content:  []string{"[STATUS: OK]", "[MENU]", "[X]"},
				Style:    "right-aligned",
				Actions: []UIAction{
					{Type: "click", Trigger: "mouse_click", Description: "Toggle status display", Handler: "toggleStatus()"},
					{Type: "click", Trigger: "mouse_click", Description: "Open menu", Handler: "openMenu()"},
					{Type: "click", Trigger: "mouse_click", Description: "Close application", Handler: "closeApp()"},
				},
			},
			{
				Type:     "separator",
				Position: Position{Y: "2"},
				Size:     Size{Width: "full"},
			},
			{
				Type:     "box",
				Name:     "System Overview Panel",
				Position: Position{X: "2", Y: "4"},
				Size:     Size{Width: "25", Height: "8"},
				Content: []string{
					"Title: SYSTEM OVERVIEW",
					"Server: PROD-WEB-01",
					"Uptime: 24d 18h 42m",
					"Processes: 127 running",
					"Users: 42 active",
				},
			},
			{
				Type:     "panel",
				Name:     "Performance Metrics",
				Position: Position{X: "30", Y: "4"},
				Size:     Size{Width: "48", Height: "8"},
				Children: []UIElement{
					{Type: "progressbar", Content: []string{"CPU Usage: 78%"}},
					{Type: "progressbar", Content: []string{"Memory: 62%"}},
					{Type: "progressbar", Content: []string{"Disk I/O: 31%"}},
					{Type: "progressbar", Content: []string{"Network: 52%"}},
				},
			},
			{
				Type:     "box",
				Name:     "Log Monitor",
				Position: Position{X: "2", Y: "13"},
				Size:     Size{Width: "76", Height: "8"},
				Content: []string{
					"Title: LOG MONITOR",
					"[14:32:01] INFO  Application started successfully",
					"[14:32:15] WARN  High memory usage detected (62%)",
					"[14:32:28] INFO  Database connection established",
					"[14:32:45] ERROR Failed to connect to backup server",
					"[14:33:02] INFO  Retrying connection...",
				},
			},
			{
				Type:     "menu",
				Position: Position{X: "2", Y: "22"},
				Content:  []string{"[Refresh]", "[Export]", "[Settings]", "[Help]"},
				Style:    "left",
				Actions: []UIAction{
					{Type: "click", Trigger: "mouse_click", Description: "Refresh data", Handler: "refreshData()"},
					{Type: "click", Trigger: "mouse_click", Description: "Export logs", Handler: "exportLogs()"},
					{Type: "click", Trigger: "mouse_click", Description: "Open settings", Handler: "openSettings()"},
					{Type: "click", Trigger: "mouse_click", Description: "Show help", Handler: "showHelp()"},
				},
			},
			{
				Type:     "statusbar",
				Position: Position{Y: "2"},
				Content:  []string{"", "Last Update: 14:33:05"},
			},
		},
	}
	
	layouts["login"] = &Layout{
		Name:        "Login Form",
		Width:       60,
		Height:      15,
		Description: "A simple login form interface",
		Elements: []UIElement{
			{
				Type:     "header",
				Name:     "Login Title",
				Position: Position{X: "center", Y: "1"},
				Content:  []string{"LOGIN SYSTEM"},
				Style:    "centered",
			},
			{
				Type:     "separator",
				Position: Position{Y: "2"},
				Size:     Size{Width: "full"},
			},
			{
				Type:     "box",
				Name:     "Login Form Container",
				Position: Position{X: "15", Y: "5"},
				Size:     Size{Width: "30", Height: "7"},
				Content: []string{
					"Username: [____________]",
					"Password: [************]",
					"",
					"[  Login  ] [ Cancel ]",
				},
				Inputs: []UIInput{
					{Type: "text", Name: "username", Validation: "alphanumeric", Required: true, Placeholder: "Enter username"},
					{Type: "password", Name: "password", Validation: "min_length:8", Required: true, Placeholder: "Enter password"},
				},
				Actions: []UIAction{
					{Type: "submit", Trigger: "click", Description: "Submit login form", Handler: "login()"},
					{Type: "reset", Trigger: "click", Description: "Cancel login", Handler: "cancel()"},
				},
			},
		},
	}
	
	layouts["network"] = &Layout{
		Name:        "Network Operations Center",
		Width:       80,
		Height:      30,
		Description: "Real-time network monitoring and operations dashboard",
		Elements: []UIElement{
			{
				Type:     "header",
				Name:     "Main Title",
				Position: Position{X: "2", Y: "1"},
				Content:  []string{"NetOps Command Center"},
			},
			{
				Type:     "header",
				Name:     "Status Bar",
				Position: Position{X: "right-25", Y: "1"},
				Content:  []string{"[LIVE]", "[ALERT]", "[X]"},
				Style:    "right-aligned",
			},
			{
				Type:     "separator",
				Position: Position{Y: "2"},
				Size:     Size{Width: "full"},
			},
			{
				Type:     "box",
				Name:     "Status Grid",
				Position: Position{X: "2", Y: "4"},
				Size:     Size{Width: "76", Height: "4"},
				Content: []string{
					"Create 6 status boxes in a row:",
					"STATUS/CORE | CPU/73% | MEMORY/12GB | NETWORK/89% | STORAGE/2.1TB | LATENCY/12ms",
				},
			},
			{
				Type:     "panel",
				Name:     "Live Traffic Flow",
				Position: Position{X: "2", Y: "9"},
				Size:     Size{Width: "30", Height: "8"},
				Content: []string{
					"Title: LIVE TRAFFIC FLOW",
					"US-WEST [====] 892 Mbps",
					"US-EAST [===] 634 Mbps",
					"EU-WEST [==] 423 Mbps",
				},
			},
			{
				Type:     "panel",
				Name:     "Incident Response",
				Position: Position{X: "34", Y: "9"},
				Size:     Size{Width: "44", Height: "8"},
				Content: []string{
					"Title: INCIDENT RESPONSE",
					"HIGH: SSL Cert Expiring 3d",
					"MED: High mem usage DB-02",
					"LOW: CDN cache miss >15%",
				},
			},
		},
	}
	
	return layouts
}

func main() {
	var (
		mode       = flag.String("mode", "list", "Mode: list, instructions, analyze")
		layout     = flag.String("layout", "dashboard", "Layout type: dashboard, login, network, custom")
		inputFile  = flag.String("input", "", "Input ASCII art file to analyze")
		outputFile = flag.String("output", "", "Output file for instructions (default: stdout)")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ASCII UI Instructor - Generate instructions for creating ASCII art UIs\n\n")
		fmt.Fprintf(os.Stderr, "This tool generates step-by-step instructions for creating ASCII art UIs\n")
		fmt.Fprintf(os.Stderr, "from UI element descriptions.\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nModes:\n")
		fmt.Fprintf(os.Stderr, "  list         - Output UI elements list with actions/inputs (default)\n")
		fmt.Fprintf(os.Stderr, "  instructions - Generate human-readable ASCII art creation instructions\n")
		fmt.Fprintf(os.Stderr, "  analyze      - Analyze existing ASCII art and output elements list\n")
		fmt.Fprintf(os.Stderr, "\nLayouts:\n")
		fmt.Fprintf(os.Stderr, "  dashboard - System monitoring dashboard\n")
		fmt.Fprintf(os.Stderr, "  login     - Simple login form\n")
		fmt.Fprintf(os.Stderr, "  network   - Network operations center\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -layout dashboard\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -mode analyze -input sample.txt\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -layout network -output instructions.md\n", os.Args[0])
	}
	
	flag.Parse()
	
	var instructions []string
	
	switch *mode {
	case "list":
		layouts := createSampleLayouts()
		layout, exists := layouts[*layout]
		if !exists {
			fmt.Fprintf(os.Stderr, "Unknown layout: %s\n", *layout)
			fmt.Fprintf(os.Stderr, "Available layouts: dashboard, login, network\n")
			os.Exit(1)
		}
		
		instructions = generateSilentOutput(layout)
		
	case "instructions":
		layouts := createSampleLayouts()
		layout, exists := layouts[*layout]
		if !exists {
			fmt.Fprintf(os.Stderr, "Unknown layout: %s\n", *layout)
			fmt.Fprintf(os.Stderr, "Available layouts: dashboard, login, network\n")
			os.Exit(1)
		}
		
		generator := NewInstructionGenerator(layout)
		instructions = generator.Generate()
		
	case "analyze":
		if *inputFile == "" {
			fmt.Fprintf(os.Stderr, "Error: -input flag is required for analyze mode\n")
			os.Exit(1)
		}
		
		content, err := os.ReadFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		
		layout := analyzeASCIIArt(string(content))
		instructions = generateSilentOutput(layout)
		
	default:
		fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", *mode)
		os.Exit(1)
	}
	
	output := strings.Join(instructions, "\n")
	
	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Instructions written to: %s\n", *outputFile)
	} else {
		fmt.Println(output)
	}
}