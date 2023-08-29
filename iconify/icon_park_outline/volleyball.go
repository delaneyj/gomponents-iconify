package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M31.81 42.405c10.168-4.319 14.907-16.056 10.598-26.213C38.098 6.035 26.351 1.276 16.193 5.595C6.035 9.914 1.275 21.65 5.595 31.808C9.914 41.965 21.652 46.724 31.81 42.405Z"/><path stroke-linecap="round" d="M16 6c-1.494 7.01 1.937 14.197 8 18M12 40c6.97-2.26 11.74-8.68 12-16m20 1c-5.45-4.672-14.5-4.597-20-1"/><path d="M17 16s11.56-8.49 24-2M20 33S7.59 28.02 7 14m27 8s1.56 14.5-10.28 22.03"/></g>`),
		g.Group(children),
	)
}