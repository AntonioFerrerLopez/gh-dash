# gh-dash

✨ A GitHub (`gh`) CLI extension to display a dashboard with **pull requests** and **issues** by filters you care about.

<img width="800px" src="https://raw.githubusercontent.com/dlvhdr/gh-prs/main/demo.gif" />

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation)
   
   _Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions._

2. Install this extension:

   ```sh
   gh extension install dlvhdr/gh-dash
   ```

3. To get the icons to render properly you should download and install a Nerd font from https://www.nerdfonts.com/.
   Then, select that font as your font for the terminal.

<details>
   <summary><strong>Installing Manually</strong></summary>

> If you want to install this extension **manually**, follow these steps:

1. Clone the repo

   ```bash
   # git
   git clone https://github.com/dlvhdr/gh-dash

   # GitHub CLI
   gh repo clone dlvhdr/gh-dash
   ```

2. Cd into it

   ```bash
   cd gh-dash
   ```

3. Install it locally
   ```bash
   gh extension install .
   ```
</details>

<details>
   <summary><strong>🌈 How do I get these exact colors and font?</strong></summary>
   
   > I'm using [Alacritty](https://github.com/alacritty/alacritty) with the [tokyonight theme](https://github.com/folke/tokyonight.nvim) and the [Fira Code](https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts/FiraCode) Nerd Font.
   > For my full setup check out [my dotfiles](https://github.com/dlvhdr/dotfiles/blob/main/.config/alacritty/alacritty.yml).
</details>

## Configuring

Configuration is provided within a `config.yml` file under the extension's directory (usually `~/.config/gh-dash/`)

The default `config.yml` file contains:

```yml
prSections:
  - title: My Pull Requests
    filters: is:open author:@me
  - title: Needs My Review
    filters: is:open review-requested:@me
  - title: Subscribed
    filters: is:open -author:@me repo:cli/cli repo:dlvhdr/gh-dash
    limit: 50 # optional limit per section
issuesSections:
  - title: Created
    filters: is:open author:@me
  - title: Assigned
    filters: is:open assignee:@me
  - title: Subscribed
    filters: is:open -author:@me repo:microsoft/vscode repo:dlvhdr/gh-dash
defaults:
  prsLimit: 20 # global limit
  issuesLimit: 20 # global limit
  preview:
    open: true
    width: 60
```

Adding a PR or issue section is as easy as adding to the list of `prSections` or `issueSections` respectively:

- title - shown in the TUI
- filters - how the repo's PRs should be filtered - these are plain [github filters](https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests)

## Usage

Run:

```sh
gh dash
```

Then press <kbd>?</kbd> for help.

## Author

Dolev Hadar dolevc2@gmail.com
