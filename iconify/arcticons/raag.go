package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="19.558" cy="36.956" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.627" ry="6.544"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.185 36.956V4.5s13.008 5.782 7.542 22.015"/>`),
		g.Group(children),
	)
}