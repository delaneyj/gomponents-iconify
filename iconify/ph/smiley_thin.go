package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M171.46 154c-9.55 16.52-25.39 26-43.46 26s-33.91-9.48-43.46-26a4 4 0 0 1 6.92-4c8.21 14.19 21.19 22 36.54 22s28.33-7.81 36.54-22a4 4 0 1 1 6.92 4ZM228 128A100 100 0 1 1 128 28a100.11 100.11 0 0 1 100 100Zm-8 0a92 92 0 1 0-92 92a92.1 92.1 0 0 0 92-92ZM92 116a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm72-16a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}