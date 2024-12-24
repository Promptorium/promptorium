## 0.1.7

Breaking changes:
- Drastic changes to config file format, see [the website](www.promptorium.org) for more info
- Removed the `--theme-file` flag, it is not needed with new config file format

New features:
- Added support for yaml config files

Config file format changes:
- New config file format:
    - `prompt` - array of strings, representing the prompt. Each string can be a component name or a module name (New in 0.1.7)
        - The prompt can also be defined as an array of arrays, representing a multiline prompt
        - `---` - indicates the **spacer** (separating the left and right sides of the prompt)
            - Each line of a multiline prompt can only contain one **spacer**
    - `theme` - a map of theme settings (already supported in 0.1.6)
    - `components` - a list of components (already supported in 0.1.6)
    - `options` - a map of options (New in 0.1.7)

- All the individual configuration fields can now be loaded from an external file, see [the documentation](www.promptorium.org) for more info
- Components now have a `type` field, which is used to determine the type of the component (module, plugin, text)
- Component dividers are no longer displayed if the background color is set to `transparent` 

Performance improvements:
- Context data is now loaded in parallel, which leads to much faster prompt rendering (now averages at ~ 10 ms)
