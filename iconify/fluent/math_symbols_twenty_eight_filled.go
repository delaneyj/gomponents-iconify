package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathSymbolsTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 2.5a1 1 0 1 0-2 0V6H3a1 1 0 0 0 0 2h3.5v3.5a1 1 0 1 0 2 0V8H12a1 1 0 1 0 0-2H8.5V2.5ZM16 6a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9Zm0 13a1 1 0 1 0 0 2h9a1 1 0 1 0 0-2h-9Zm4.75-1a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 5.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM2.293 15.293a1 1 0 0 1 1.414 0L7.5 19.086l3.793-3.793a1 1 0 1 1 1.414 1.414L8.914 20.5l3.793 3.793a1 1 0 0 1-1.414 1.414L7.5 21.914l-3.793 3.793a1 1 0 0 1-1.414-1.414L6.086 20.5l-3.793-3.793a1 1 0 0 1 0-1.414Z"/>`),
		g.Group(children),
	)
}