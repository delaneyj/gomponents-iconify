package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 8a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1H6Zm2-7a2 2 0 0 0-2 2v.416A5.001 5.001 0 0 0 3 8v2h10V8a5.001 5.001 0 0 0-3-4.584V3a2 2 0 0 0-2-2Zm0 2c-.342 0-.677.034-1 .1V3a1 1 0 0 1 2 0v.1A5.023 5.023 0 0 0 8 3ZM7 6h2a2 2 0 0 1 2 2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1a2 2 0 0 1 2-2Zm-1 6.5V11H3v1.5A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2.5-2.5V11H7v1.5a.5.5 0 0 1-1 0Z"/>`),
		g.Group(children),
	)
}