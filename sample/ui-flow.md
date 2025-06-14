# UI Flow Diagram

```mermaid
flowchart TB
    Start([Start])
    login[Login<br/>login.txt]
    main[Main<br/>main.txt]
    reports[Reports<br/>reports.txt]
    settings[Settings<br/>settings.txt]
    Start --> login
    login -->|[Login]\nSubmit login| main
    main -->|[Logout]\nLogout| login
    main -->|[View Reports]\nOpen reports| reports
    main -->|[Settings]\nOpen settings| settings
    reports -->|[Back]\nReturn to main| main
    settings -->|[Cancel]\nCancel changes| main
```

## Screen Transition Details

Actions available on each screen and their destinations:

### login.txt
- **Username**: input: text input → (no screen transition)
- **Password**: input: password input → (no screen transition)
- **[Login]**: click: Submit login → main.txt
- **[Exit]**: click: Exit application → (application exit)

### main.txt
- **[Logout]**: click: Logout → login.txt
- **[View Reports]**: click: Open reports → reports.txt
- **[Settings]**: click: Open settings → settings.txt
- **[Help]**: click: Show help → (no screen transition)

### reports.txt
- **[Back]**: click: Return to main → main.txt
- **[Daily Summary]**: click: View daily report → (no screen transition)
- **[Weekly Analysis]**: click: View weekly report → (no screen transition)
- **[Monthly Report]**: click: View monthly report → (no screen transition)
- **[Custom Report]**: click: Create custom report → (no screen transition)
- **[Download]**: click: Download selected report → (no screen transition)
- **[View]**: click: View selected report → (no screen transition)
- **[Delete]**: click: Delete selected report → (no screen transition)
- **[Generate New]**: click: Generate new report → (no screen transition)

### settings.txt
- **[Back]**: click: Return to main → main.txt
- **> General**: click: General settings → (no screen transition)
- **> Security**: click: Security settings → (no screen transition)
- **> Notifications**: click: Notification settings → (no screen transition)
- **[Save Changes]**: click: Save all settings → main.txt
- **[Reset]**: click: Reset to defaults → (no screen transition)
- **[Cancel]**: click: Cancel changes → main.txt

## Usage Guide

1. When creating ASCII UI for each screen, place the UI at the top
2. Add the `INTERACTIVE_ELEMENTS:` section at the bottom of the screen
3. Describe each interactive element in the following format:
   ```
   [Element Name] - action_type: Description → destination.txt
   ```
   - For no transition: `→ (no screen transition)`
   - For app exit: `→ (application exit)`
