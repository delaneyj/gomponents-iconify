package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XboxConsoleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 2a.75.75 0 0 0-.75.75v18.5c0 .414.336.75.75.75H8V11.75a.75.75 0 0 1 1.5 0V22h8.75a.75.75 0 0 0 .75-.75V2.75a.75.75 0 0 0-.75-.75H5.75ZM9.5 5.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}