package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VietinbankIpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.5 30.483V17.5h13v12.983"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.15 34.341a43 43 0 0 1 37.7 0"/>`),
		g.Group(children),
	)
}