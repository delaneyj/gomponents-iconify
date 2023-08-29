package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommandBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8h3a3 3 0 1 0-3-3v3H8V5a3 3 0 1 0-3 3h3v8h8V8Zm0 8h3a3 3 0 1 1-3 3.001V16ZM5 16l3 .001v3a3 3 0 1 1-3-3Z"/>`),
		g.Group(children),
	)
}