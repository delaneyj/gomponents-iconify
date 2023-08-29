package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Groupme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.9 6.594v19.377m8.2-19.377v19.377m5.589-13.788H14.312m19.377 8.2H14.312M4.5 22.122a19.501 19.501 0 0 0 18.403 19.254A19.501 19.501 0 0 0 43.5 22.122"/>`),
		g.Group(children),
	)
}