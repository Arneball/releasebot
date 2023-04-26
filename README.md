# 🤖 Releasebot

> **Because Fastlane is slow, overcomplicated, and I'm tired of waiting for Ruby to do simple things.**

A blazingly fast CLI tool written in Go for Android release management and Teams notifications. Does what Fastlane does, but without the headaches.

## Why This Exists

Fastlane is great in theory, but in practice:
- ⏰ **Slow**: Ruby startup time + dependency loading = eternal waiting
- 🐌 **Bloated**: Need 50 gems just to upload a file?
- 🔧 **Overengineered**: Simple tasks require complex configuration
- 😤 **Frustrating**: When you just want to upload an APK and move on with your life

**Releasebot** is the answer: a single binary that does exactly what you need, nothing more.

## Installation

### Build from source
```bash
go build -o releasebot
```

### Deploy to remote servers
```bash
cd ansible
ansible-playbook -i hosts.yaml playbook.yml
```

## Setup

Set the path to your Google Play service account JSON credentials:

```bash
export FASTLANE_KEY=/path/to/your/service-account.json
```

Or add it to your `.env` file (gitignored):
```
FASTLANE_KEY=/path/to/your/service-account.json
```

## Typical Usage

Here's how I actually use it in my workflow - one command to build and upload:

```bash
VERSION_CODE=$(releasebot nextVersionCode --package=com.vgregion.hud) \
  ./gradlew hud:bundleRelease && \
  releasebot upload \
    --package=com.vgregion.hud \
    --track=internal \
    --aab=hud/build/outputs/bundle/release/*.aab \
    --status=completed
```

This fetches the next version code, builds the release bundle with that version, and immediately uploads it to the internal track. No waiting, no Ruby, just done.

## Commands

### 🚀 Upload AAB to Google Play

Upload an Android App Bundle to a specific track:

```bash
releasebot upload \
  --package com.example.app \
  --track internal \
  --aab ./app-release.aab \
  --status completed
```

**Options:**
- `--package`: Android package name (e.g., `com.vgregion.hud`)
- `--track`: Release track (`internal`, `alpha`, `beta`, `production`)
- `--aab`: Path to your .aab file
- `--status`: `draft` or `completed`

### 🔢 Get Next Version Code

Find out what version code to use for your next build:

```bash
releasebot nextVersionCode --package com.example.app
```

Returns the next available version code (e.g., `42`).

### 📦 Update Release Track

Promote an existing version to a different track:

```bash
releasebot updateTrack \
  --package com.example.app \
  --version 42 \
  --track production
```

Perfect for promoting an internal release to beta or production.

### 📲 Download APK

Download a universal APK from Google Play Console:

```bash
releasebot downloadApk \
  --package com.example.app \
  --version 42
```

Downloads to `/tmp/com.example.app_42.apk`.

### 💬 Send Teams Notification

Send a message to Microsoft Teams with a QR code:

```bash
echo "Build complete! 🎉" | releasebot teams-notify \
  --title "Download APK" \
  --url "https://example.com/download" \
  --webhook_url "https://outlook.office.com/webhook/..."
```

Automatically generates a QR code for the URL and includes it in the Teams message.

## Shell Completion

Releasebot includes smart shell completion for common values:

```bash
# Tab completion works for package names
releasebot upload --package <TAB>
# Suggests: com.vgregion.hud, com.vgregion.migraine, com.vgregion.epilepsy

# Tab completion for .aab files
releasebot upload --package com.example.app --track internal --aab <TAB>
# Lists all .aab files in current directory

# Tab completion for status
releasebot upload --package com.example.app --track internal --aab app.aab --status <TAB>
# Suggests: draft, completed
```

## Architecture

Built with:
- **[Cobra](https://github.com/spf13/cobra)**: Clean CLI interface
- **[Google Play Developer API](https://developers.google.com/android-publisher)**: Direct API access, no middleman
- **[go-teams-notify](https://github.com/atc0005/go-teams-notify)**: Microsoft Teams integration
- **[go-qrcode](https://github.com/skip2/go-qrcode)**: QR code generation

Two main packages:
- **`slowlane/`**: Google Play Console operations (upload, download, version management)
- **`botstuff/`**: Teams notifications with QR codes

## Development

### Run tests
```bash
go test ./...
```

### Add a new package to completions

Edit `slowlane/download_apk.go` and add your package to the `ourPackages()` function:

```go
func ourPackages(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
    pkgs := []string{
        "com.vgregion.hud",
        "com.vgregion.migraine",
        "com.vgregion.epilepsy",
        "com.your.new.package", // Add here
    }
    return pkgs, cobra.ShellCompDirectiveNoSpace
}
```

## License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0) - see the [LICENSE](LICENSE) file for details.

This means you're free to use, modify, and distribute this software, but if you run a modified version as a service or distribute it, you must release your source code under the same license. No one gets to steal this and make it proprietary.

**Attribution required**: If you use this, give credit. I'm an attention whore and I want people to know who saved them from Fastlane hell.

---

*Made with ❤️ (and frustration) by someone who just wanted to ship apps faster.*
