package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8H8.001L8 20H6V8H2l5-5l5 5Zm10 8l-5 5l-5-5h4V4h2v12h4Z"/>`),
		g.Group(children),
	)
}