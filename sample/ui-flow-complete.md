# UI Flow Diagram

```mermaid
flowchart TB
    Start([Start])
    dashboard[Dashboard<br/>dashboard.txt]
    login[Login<br/>login.txt]
    sample[Sample<br/>sample.txt]
    Start --> login
    dashboard --> login
    dashboard --> netops
    dashboard --> trading
    dashboard --> cyber
    dashboard --> data
    dashboard --> space
    dashboard --> settings
    dashboard --> help
    login --> dashboard
    sample --> menu_dropdown
    sample --> export_dialog
    sample --> settings
    sample --> help
```

## Detected Screens

- **dashboard.txt** → login.txt, netops.txt, trading.txt, cyber.txt, data.txt, space.txt, settings.txt, help.txt
- **login.txt** → dashboard.txt
- **sample.txt** → menu_dropdown.txt, export_dialog.txt, settings.txt, help.txt
