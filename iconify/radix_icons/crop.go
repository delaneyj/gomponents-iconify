package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 0a.5.5 0 0 1 .5.5V3h7.5a.5.5 0 0 1 .5.5V11h2.5a.5.5 0 1 1 0 1H12v2.5a.5.5 0 0 1-1 0V12H3.5a.5.5 0 0 1-.5-.5V4H.5a.5.5 0 1 1 0-1H3V.5a.5.5 0 0 1 .5-.5ZM4 4v7h7V4H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}