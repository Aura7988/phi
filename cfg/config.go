package cfg

import "strconv"

type TomlConfig struct {
	Editor EditorConfig
	Cursor CursorConfig
	Render RenderConfig
	Theme  ThemeConfig
}

var DEFUALT_TOML_CONFIG string = `[editor]
tab_size = 2
hungry_backspace = true
tabs_are_spaces = true
match_braces = false
maintain_indentation = true
highlight_line = true

[render]
aliased = true
accelerated = true

[theme]
background = "0x002649"
foreground = "0xf2f4f6"
cursor = "0xf2f4f6"
cursor_invert = "0xffffff"

[cursor]
flash_rate = 400
reset_delay = 400
draw = true
flash = true
`

type CursorConfig struct {
	Flash_Rate  int64
	Reset_Delay int64
	Draw        bool
	Flash       bool
	Block_Width string
}

func (c CursorConfig) GetCaretWidth() int {
	if c.Block_Width == "block" {
		return -1
	}
	if c.Block_Width == "" {
		return -1
	}

	value, err := strconv.ParseInt(c.Block_Width, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(value)
}

type RenderConfig struct {
	Aliased     bool
	Accelerated bool
}

// todo make this more extendable...
// e.g. .phi-editor/themes with TOML
// themes in them and we can select
// the default theme in the EditorConfig
// instead.
type ThemeConfig struct {
	Background    string
	Foreground    string
	Cursor        string
	Cursor_Invert string
}

type EditorConfig struct {
	Tab_Size             int
	Hungry_Backspace     bool
	Tabs_Are_Spaces      bool
	Match_Braces         bool
	Maintain_Indentation bool
	Highlight_Line       bool
}

func NewDefaultConfig() *TomlConfig {
	return &TomlConfig{
		Editor: EditorConfig{},
		Theme: ThemeConfig{
			Background:    "0x002649",
			Foreground:    "0xf2f4f6",
			Cursor:        "0xf2f4f6",
			Cursor_Invert: "0xffffff",
		},
	}
}
