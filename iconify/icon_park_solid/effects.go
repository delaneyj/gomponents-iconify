package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Effects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEffects0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m24.95 42.36l5.466-11.99l12.689-3.72l-9.767-8.88l.368-13.163l-11.502 6.503l-12.46-4.416l2.657 12.9l-8.069 10.433l13.145 1.47L24.95 42.36Z"/><path d="m36.178 36.054l8 7.964"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEffects0)"/>`),
		g.Group(children),
	)
}