package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.15 14V1H.85v13h1.3Zm4 0V1h-1.3v13h1.3Zm4-13v13h-1.3V1h1.3Zm4 13V1h-1.3v13h1.3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}