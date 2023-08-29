package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunicationTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 4a7 7 0 0 0-4.952 11.948a.5.5 0 0 1-.707.707a8 8 0 1 1 11.318 0a.5.5 0 0 1-.707-.707A7 7 0 0 0 10 4Zm0 3a4 4 0 0 0-2.832 6.825a.5.5 0 1 1-.708.706a5 5 0 1 1 7.078.002a.5.5 0 0 1-.708-.706A4 4 0 0 0 10 7Zm0 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-1 2a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		g.Group(children),
	)
}