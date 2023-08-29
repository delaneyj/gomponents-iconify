package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MetronomeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 216H56a8 8 0 0 1-7.63-10.43l12-37.57h135.29l12 37.57A8 8 0 0 1 200 216Z" opacity=".2"/><path d="m187.14 114.84l26.78-29.46a8 8 0 0 0-11.84-10.76l-20.55 22.6l-17.2-54.07A15.94 15.94 0 0 0 149.08 32h-42.17a15.94 15.94 0 0 0-15.25 11.15l-50.91 160A16 16 0 0 0 56 224h144a16 16 0 0 0 15.25-20.85ZM184.72 160h-38.64l28.62-31.48ZM106.91 48h42.17l20 62.9l-44.62 49.1H71.27ZM56 208l10.18-32h123.63L200 208Z"/></g>`),
		g.Group(children),
	)
}