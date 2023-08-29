package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotistats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="7.68" height="34.586" x="20.16" y="6.707" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.327"/><rect width="7.68" height="15.244" x="33.821" y="26.049" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.327"/><rect width="7.68" height="24.721" x="6.5" y="16.572" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.327"/>`),
		g.Group(children),
	)
}