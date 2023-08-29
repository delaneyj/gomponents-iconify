package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 9l-2-4h-3V3c0-.55-.45-1-1-1H1c-.55 0-1 .45-1 1v8l1 1h1.268a2 2 0 1 0 3.464 0h5.536a2 2 0 1 0 3.464 0H16V9zm-5 0V6h2.073l1.5 3H11z"/>`),
		g.Group(children),
	)
}