package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LedgerLive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.797 39.203A21.504 21.504 0 0 1 32.734 4.348m6.469 4.449a21.504 21.504 0 0 1-23.937 34.855m-6.469-4.449l5.657-5.657m19.092-19.092l5.657-5.657"/>`),
		g.Group(children),
	)
}