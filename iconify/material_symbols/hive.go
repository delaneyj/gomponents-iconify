package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.275 11.5l-1.7-3l1.7-3h3.35l1.7 3l-1.7 3h-3.35Zm-5.95 3.5l-1.7-3l1.7-3h3.35l1.7 3l-1.7 3h-3.35Zm0-7l-1.7-3l1.7-3h3.35l1.7 3l-1.7 3h-3.35Zm-5.95 3.5l-1.7-3l1.7-3h3.35l1.625 3l-1.625 3h-3.35Zm0 7l-1.7-3l1.7-3h3.35l1.625 3l-1.625 3h-3.35Zm6.05 3.5l-1.8-3l1.7-3h3.35l1.7 3l-1.7 3h-3.25Zm5.85-3.5l-1.7-3l1.7-3h3.35l1.7 3l-1.7 3h-3.35Z"/>`),
		g.Group(children),
	)
}