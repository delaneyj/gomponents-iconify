package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.62 1024q-53 0-90.5-37.5t-37.5-90.5V128q0-53 37.5-90.5T896.62 0t90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-834-11q-25 27-62-13V24q37-40 62-13l563 463q15 16 15 38.5t-15 37.5z"/>`),
		g.Group(children),
	)
}