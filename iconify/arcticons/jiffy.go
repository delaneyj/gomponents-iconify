package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jiffy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.429 2.967A21.504 21.504 0 0 0 24 2.5a21.501 21.501 0 0 0-4.389.508m24.814 14.408a21.502 21.502 0 0 0-13.729-13.84m-13.359.022a21.363 21.363 0 0 0-14.326 24.99M30.674 44.42A21.5 21.5 0 0 0 45.5 24a21.506 21.506 0 0 0-.49-4.31h0M3.643 30.846a21.411 21.411 0 0 0 24.763 14.186M17.797 13.114l8.356 15.762m-.136-6.166l10.9-6.572M19.628 26.687l2.134-1.374"/>`),
		g.Group(children),
	)
}