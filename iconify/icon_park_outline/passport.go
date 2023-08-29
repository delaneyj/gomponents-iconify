package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 10h28v34H10V10Zm0 0l22-6v6"/><circle cx="24" cy="24" r="4"/><path d="M20 34h8"/></g>`),
		g.Group(children),
	)
}