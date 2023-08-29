package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.17 17.74l-.71-.7L20 12.5l-4.54-4.54l.71-.7l5.24 5.24l-5.24 5.24m-9.34 0L1.59 12.5l5.24-5.24l.71.7L3 12.5l4.54 4.54l-.71.7M12.73 5h1l-3.46 15h-1l3.46-15Z"/>`),
		g.Group(children),
	)
}