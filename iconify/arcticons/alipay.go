package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alipay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2Zm-30.17 6.72h22.28M23.47 8.5v10.66"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 34c-6.65-3.9-33.57-17.8-33.57-3.78c0 3.48 3.4 5.17 6.68 5.17c7.86 0 15.25-11 15.25-16.22H16.08"/>`),
		g.Group(children),
	)
}