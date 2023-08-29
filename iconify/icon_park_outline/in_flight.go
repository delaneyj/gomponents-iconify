package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InFlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M12 40c0-7.18 5.373-13 12-13s12 5.82 12 13"/><path d="M13 37c0 1 2.5 2 4 1s1.96-3.505 3.5-3.074C22.04 35.357 22 38.5 24 40s5.5 1 7-1.5s-.08-3.175 1.061-4.797c.76-1.081 1.73-.816 1.939-.703"/><path stroke-linejoin="round" d="M23 21h2m7.151 1.47l1.696 1.06m-19.696 0l1.697-1.06m23.211 5.664l1 1.732m-32 0l1-1.732m32.836 8.872l.209 1.989m-37.209 0l.209-1.99"/><path d="M42 10L9 18"/><path stroke-linejoin="round" d="m16 7l13 6l-12 3l-5-7l4-2ZM9 18l-3-4"/></g>`),
		g.Group(children),
	)
}