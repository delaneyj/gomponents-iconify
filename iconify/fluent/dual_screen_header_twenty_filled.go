package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenHeaderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 16h-5.5V8H18v6a2 2 0 0 1-2 2Zm2-9V6a2 2 0 0 0-2-2h-5.5v3H18ZM9.5 7V4H4a2 2 0 0 0-2 2v1h7.5ZM2 8v6a2 2 0 0 0 2 2h5.5V8H2Z"/>`),
		g.Group(children),
	)
}