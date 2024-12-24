---
sidebar_position: 3
---

# Configuration

Promptorium allows you to customize it using the `config.yaml` file located by default at `~/.config/promptorium/`.

The `config.yaml` file contains the following sections:

- `prompt`
- `components`
- `theme`
- `options`

## Prompt

The `prompt` section contains an array of strings, representing the prompt. Each string can be a [module](#modules) name or a [component](#components) name prefixed with `$`.
e.g. 

```yaml title="~/.config/promptorium/config.yaml"
prompt: [ 'module_name', '$component_name' ]
components:
  ...
```

In this example, the prompt will display the `module_name` module and the `component_name` component defined in the `components` section.

The prompt can also contain the special `---` string, indicating the **spacer**.
The spacer is used to separate the left and right sides of the prompt.
e.g.
```yaml title="~/.config/promptorium/config.yaml"
prompt: [ 'module_name', '---', '$component_name' ]
components:
  ...
```

In this example, the prompt will display the `module_name` module to the left and the `component_name` component defined in the `components` section to the right, separated by a **spacer**.

The prompt can also contain an array of arrays, representing a multiline prompt.
e.g.
```yaml title="~/.config/promptorium/config.yaml"
prompt:
  - [ 'module_name', '---', '$component_name' ]
  - [ '$another_component']
```

This prompt configuration will display two lines of the prompt, the first line containing `module_name`, `---`, and `component_name`, and the second line containing `another_component`.


## Components

Components are the building blocks of the prompt. Each component has a name, content, and style. The content is the actual content of the module, while the style is used to customize the appearance of the module.

The `config.yaml` file contains an array of objects, representing components. Each object has the following structure:

```yaml title="~/.config/promptorium/config.yaml"

components:
- name: 'component_name'        //required
  type: 'component_type'        //required
  content:                      //required
  style:                        //optional
    ...

```

### Name (Required)

The `name` field is the name of the component. It is needed to uniquely identify the component in the `prompt` field.


### Type (Required)

The `type` field is the type of the component. It can be one of the following:
- `module`
- `text`

### Content (Required)

The `content` field contains the actual content of the component. Depending on the type of the component, the content will be processed differently:

- For `module` components, the content is the name of the module to be displayed in the component.

- For `text` components, the content is the text to be displayed in the component.

#### Module (Required)

The `module` field is the name of the module to be displayed in the component. Promptorium supports the following modules:

- `time`
- `hostname`
- `cwd`
- `git_branch`
- `exit_status`
- `user`
- `os_icon`
- `git_status`

For more details on each module, see the [Modules](#modules) section.

#### Icon (Optional)

The `icon` field is the character that will be displayed as the icon of the component. By default it is an empty string.

:::warning 
Promptorium needs the `icon` to be one character long. Setting anicon with more than one character will result in incorrect behavior.
:::

### Style (Optional)

The `style` field contains the style of the component. By default it is defined by the theme.

The `style` object has the following structure:

```yaml
style: 
    background_color: "background_color",
    foreground_color: "foreground_color",
    start_divider: "divider_character",
    end_divider: "divider_character",
    margin: "margin",
    padding: "padding",
    icon_background_color: "background_color",
    icon_foreground_color: "foreground_color",
    icon_padding: "padding",
    icon_position: "left|right"
```

#### Background Color (Optional)

The `background_color` field is the background color of the component. See the [Colors](#colors) section for more details on the available colors.

#### Foreground Color (Optional)

The `foreground_color` field is the foreground(text) color of the component.

#### Start Divider (Optional)

The `start_divider` field is the character that will be displayed at the beginning of the component. By default it is the theme's start divider.

#### End Divider (Optional)

The `end_divider` field is the character that will be displayed at the end of the component. By default it is the theme's end divider.

:::warning 
Promptorium needs the `start_divider` and `end_divider` to be one character long. Setting them with more than one character will result in incorrect behavior.
:::

Components aligned to the right are displayed on the right side of the prompt, separated from components aligned to the left by a string composed of [spacer](#spacer-optional) characters.
#### Margin (Optional)

The `margin` field is the margin of the component, or the space between the component and other components. It can be either a single number or two numbers separated by a space.

If it is a single number, it is applied to the left and right margin. If it is a two numbers, the first number is applied to the left margin and the second number is applied to the right margin. By default it is set to `1 0`.

#### Padding (Optional)

The `padding` field is the padding of the component, or the space between the separator and the component's content.

Similarly to the `margin` parameter, it can be set to a single number or two numbers separated by a space. By default it is set to `1`.

#### Icon Background Color (Optional)

The `icon_background_color` field is the background color of the icon. See the [Colors](#colors) section for more details on the available colors.

#### Icon Foreground Color (Optional)

The `icon_foreground_color` field is the foreground(text) color of the icon. See the [Colors](#colors) section for more details on the available colors.

#### Icon Padding (Optional)

The `icon_padding` field is the padding of the icon, or the space between the icon and the component's content.
Unlike the `padding` parameter, it can only be set to a single number. By default it is set to `1`.

#### Icon Position (Optional)

The `icon_position` field is the position of the icon relative to the component. It can be either `left` or `right`. By default it is set to `left`.

## Theme

The `theme` section contains the following fields:
```yaml title="~/.config/promptorium/config.yaml"
theme:
    component_start_divider: component_start_divider
    component_end_divider: component_end_divider
    primary_color: color
    secondary_color: color
    tertiary_color: color
    quaternary_color: color
    success_color: color
    warning_color: color
    error_color: color
    background_color: color
    foreground_color: color
    git_status_clean: git_status_clean
    git_status_dirty: git_status_dirty
    git_status_no_repository: git_status_no_repository
    git_status_no_upstream: git_status_no_upstream

```

#### Component Start Divider (Optional)

The `component_start_divider` field is the character that will be displayed at the beginning of each component unless the `start_divider` field of the component is set.

#### Component End Divider (Optional)

The `component_end_divider` field is the character that will be displayed at the end of each component unless the `end_divider` field of the component is set.

#### Primary Color (Optional)

The `primary_color` field is the primary color of the theme. Default value is "blue"

#### Secondary Color (Optional)

The `secondary_color` field is the secondary color of the theme. Default value is "green"

#### Tertiary Color (Optional)

The `tertiary_color` field is the tertiary color of the theme. Default value is "cyan"

#### Quaternary Color (Optional)

The `quaternary_color` field is the quaternary color of the theme. Default value is "magenta"

#### Success Color (Optional)

The `success_color` field is the success color of the theme. Default value is "green"

#### Warning Color (Optional)

The `warning_color` field is the warning color of the theme. Default value is "yellow"

#### Error Color (Optional)

The `error_color` field is the error color of the theme. Default value is "red"

#### Background Color (Optional)

The `background_color` field is the background color of the theme. Default value is "transparent"

#### Foreground Color (Optional)

The `foreground_color` field is the foreground(color) color of the theme. Default value is "white"

#### Git Status Clean (Optional)

The `git_status_clean` field is the color of the git status when the repository is clean. Default value is "green"

#### Git Status Dirty (Optional)

The `git_status_dirty` field is the color of the git status when the repository is dirty. Default value is "yellow"

#### Git Status No Repository (Optional)

The `git_status_no_repository` field is the color of the git status when no branch is checked out. Default value is "blue"

#### Git Status No Upstream (Optional)

The `git_status_no_upstream` field is the color of the git status when no upstream branch is checked out. Default value is "white"

## Colors

Promptorium has three types of color parameters: ***base colors***, ***theme colors*** and ***color functions***.

**[Base colors](#base-colors)** are predefined colors that you can use as-is.

**[Theme colors](#theme-colors)** are colors that you can customize using the `theme.json` file.

**[Color functions](#color-functions)** are special colors which change depending on the context of the application.


For example, here is an example of a **base color**:

```yaml

---
name: my_component
content:
  module: my_module
style:
  // highlight-next-line
  background_color: blue

```

In this case, the background color is being set to a **base color**.

Here is an example of a **theme color**:
```yaml
name: my_component
content:
  module: my_module
style:
  // highlight-next-line
  background_color: "$primary_color"
  ...
theme:
  primary_color: red

```
In this case, the background color is being set to a **theme color**. You can customize the theme colors in the `theme.json` file.

Here is an example of a **color function**:
```yaml
name: my_component
content:
  module: my_module
style:
  // highlight-next-line
  background_color: "$exit_code_color"

```

In this case, the background color is being set to a **color function**. You can customize the color function in the `theme.json` file.

You can find more information about color functions in the [Color Functions](#color-functions) section.

:::info
Theme colors and color functions can only be set to **base colors**. (e.g. `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`,`transparent`)
:::


### Base Colors

Base colors are predefined colors that you can use as-is. They are depend on your terminal emulator, so if you want to change the shade of a base color, you have to change it in the terminal emulator's settings.

Here are the available base colors:
- `black`
- `red`
- `green`
- `yellow`
- `blue`
- `magenta`
- `cyan`
- `white`
- `transparent`

### Theme Colors

Theme colors are colors that you can customize using the `theme.json` file.
Here are the available theme colors:
- `primary_color` 
- `secondary_color`
- `tertiary_color` 
- `quaternary_color`
- `success_color`
- `warning_color`
- `error_color` 
- `background_color` 
- `foreground_color` 

:::info
When using theme colors or color functions in components, remember to put a `$` in front of the color name.
:::

### Color Functions

Color functions are special colors which change depending on the context of the application.

Here are the available color functions:
- `exit_code_color`
- `git_status_color`

You can customize the color for each of the color functions' states in the `theme.json` file.

#### exit_code_color

The `exit_code_color` color function is used to display the exit status of last executed command.
Here is the color function's states and corresponding colors:
- success: the previous command returned 0. Uses the `success_color` theme color.
- error: the previous command returned a non-zero exit status. Uses the `error_color` theme color.

#### git_status_color

The `git_status_color` color function is used to display the status of the git repository.
Here is the color function's states and corresponding colors:
- clean: the repository is clean, meaning that there are no uncommitted changes. Uses the `git_status_clean` theme color.
- dirty: the repository is dirty, meaning that there are uncommitted changes. Uses the `git_status_dirty` theme color.
- no-repository: not in a git repository. Uses the `git_status_no_repository` theme color.
- no-upstream: current local branch does not have an upstream branch. Uses the `git_status_no_upstream` theme color.


## Modules

Modules are used to display different information in the promptorium prompt. They can be used in the `config.yaml` file to customize the appearance of the prompt.
Here are the available modules:

### user

The `user` module displays the current user.

### cwd

The `cwd` module displays the current working directory.

### git_branch

The `git_branch` module displays the current git branch.

### git_status

The `git_status` module displays icons representing the status of the git repository. The icons represent the following states:
- Regarding changes to the workspace:
    - `empty-circle`: Unstaged changes
    - `filled-circle`: Staged but uncommitted changes
    - `checkmark`: No changes
- Regarding differences between the local and upstream repository:
    - `arrow-up`: Ahead of upstream
    - `arrow-down`: Behind upstream
    - `arrow-up` and `arrow-down`: Diverged from upstream

### os_icon

The `os_icon` module displays the operating system icon. Not all operating systems and distributions are supported, but the `os_icon` module is always available. Here is the list of supported ones:
- Arch Linux
- Debian
- Fedora
- Ubuntu
- Mac OS

If the operating system is not supported, the `os_icon` module will display a default icon.

### time

The `time` module displays the current time, formatted as `HH:MM:SS`.

### hostname

The `hostname` module displays the current hostname.


## Presets

Presets are useful when you want to change between different promptorium configurations.
Promptorium allows you to create presets in the `~/.config/promptorium/presets/` directory.

To use a preset, set the `preset` field in the `config.yaml` file to the name of the preset you want to use. Promptorium will load the preset's `config.yaml` and `theme.json` files by searching `~/.config/promptorium/presets/` for a directory with the same name as the preset.
Here is an example :

Let's say you have two presets, `default_1` and `default_2`. The `config.yaml` and `theme.json` files for each preset are in the following folder structure:
```bash
~/.config/promptorium/
├── presets
│   ├── default-1
│   │   ├── config.yaml
│   │   ├── components.yaml
│   │   └── theme.yaml
│   └── default-2
│       ├── config.yaml
│       ├── components.yaml
│       └── theme.yaml
└── config.yaml
```

To use the `default-1` preset, set the `preset` key in the `config.yaml` file to `default-1`. Promptorium will load the `default-1` preset's `config.yaml` file.

```yaml title="~/.config/promptorium/config.yaml"
preset : default-1

```

:::info
If the `preset` key in the `config.yaml` file is set, Promptorium will ignore the rest of the `conf.yaml` file.
:::
 


## Options

Options are used to configure the promptorium prompt. They can be used in the `config.yaml` file to customize the appearance of the prompt.

Here are the available option categories and their corresponding options:

### cwd

The `cwd` option is used to configure the cwd module.

- `highlight_git_root` (bool): If true, the root of the git repository will be underlined in the cwd module. Default value is false. In cases of nested git repositories, the root of the outermost repository will be underlined.