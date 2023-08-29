package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digipost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.292 5.106h31.147v39H9.292zm0 10.411h31.147"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.528 15.517l-15-7.676l-4.059 7.676m3.132 6.618h12.397v4.632H18.601zm1.809 2.316h6.397"/>`),
		g.Group(children),
	)
}