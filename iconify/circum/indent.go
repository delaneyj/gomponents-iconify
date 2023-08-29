package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.437 4.064H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1Zm0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 5.624H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1ZM7.91 11.65a.492.492 0 0 1 0 .7l-2 2a.495.495 0 0 1-.7-.7l1.15-1.15H3.54a.5.5 0 0 1 0-1h2.81c-.38-.38-.76-.76-1.14-1.15a.495.495 0 0 1 .7-.7Z"/>`),
		g.Group(children),
	)
}