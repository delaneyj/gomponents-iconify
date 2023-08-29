package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h2v16H2V4Zm4 0h2v16H6V4Zm3 0h3v16H9V4Zm4 0h2v16h-2V4Zm3 0h2v16h-2V4Zm3 0h3v16h-3V4Z"/>`),
		g.Group(children),
	)
}