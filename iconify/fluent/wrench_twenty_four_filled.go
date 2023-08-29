package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.75 2.001a5.25 5.25 0 0 0-5.005 6.839l-9.068 9.38a2.344 2.344 0 1 0 3.37 3.258l8.963-9.272a5.25 5.25 0 0 0 6.786-6.407a.75.75 0 0 0-1.251-.323L17.361 8.66l-2.06-2.06l3.16-3.162a.75.75 0 0 0-.333-1.254A5.255 5.255 0 0 0 16.75 2Z"/>`),
		g.Group(children),
	)
}