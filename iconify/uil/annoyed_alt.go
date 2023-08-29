package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnnoyedAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.66 13.56l-4.19 1.5A1 1 0 0 0 10.8 17a1 1 0 0 0 .34-.06l4.2-1.5a1 1 0 1 0-.68-1.88Zm-4-5a1 1 0 0 0-1.41 0a1 1 0 0 1-1.42 0a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42a3 3 0 0 0 4.24 0a1 1 0 0 0-.04-1.44Zm7 0a1 1 0 0 0-1.41 0a1 1 0 0 1-1.42 0a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42a3 3 0 0 0 4.24 0a1 1 0 0 0-.04-1.44ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}