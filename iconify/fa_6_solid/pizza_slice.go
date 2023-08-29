package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PizzaSlice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M169.7.9c-22.8-1.6-41.9 14-47.5 34.7L110.4 80h1.6c176.7 0 320 143.3 320 320v1.6l44.4-11.8c20.8-5.5 36.3-24.7 34.7-47.5C498.5 159.5 352.5 13.5 169.7.9zm230.1 409.3c.1-3.4.2-6.8.2-10.2c0-159.1-128.9-288-288-288c-3.4 0-6.8.1-10.2.2L.5 491.9c-1.5 5.5.1 11.4 4.1 15.4s9.9 5.6 15.4 4.1l379.8-101.2zM176 208a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm64 128a32 32 0 1 1 64 0a32 32 0 1 1-64 0zM96 384a32 32 0 1 1 64 0a32 32 0 1 1-64 0z"/>`),
		g.Group(children),
	)
}