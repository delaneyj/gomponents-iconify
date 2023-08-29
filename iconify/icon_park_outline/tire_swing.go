package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TireSwing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 4s6 6 20 6s20-6 20-6m-20 6v6"/><ellipse cx="30" cy="30" rx="8" ry="14"/><ellipse cx="30" cy="30" rx="3" ry="6"/><path d="M18 44c-4.418 0-8-6.268-8-14s3.582-14 8-14m12 0H18m12 28H18m4-15H10m13-7H12m11 15H12"/></g>`),
		g.Group(children),
	)
}