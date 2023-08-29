package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DustTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 3a6.502 6.502 0 0 0-6.258 4.736a6.5 6.5 0 1 0 0 12.527A6.5 6.5 0 1 0 21.19 14A6.5 6.5 0 0 0 16.5 3Zm-4.955 5.829a5.001 5.001 0 1 1 8.068 4.584a.75.75 0 0 0 0 1.174a5 5 0 1 1-8.068 4.585a.75.75 0 0 0-1.044-.588a5 5 0 1 1 0-9.168a.75.75 0 0 0 1.044-.587ZM9 4a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm16 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM8 25a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}