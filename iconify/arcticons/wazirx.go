package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wazirx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.669 36.444L43.5 13.81L14.612 36.043m-8.225 0L24 5.884m-8.884 15.213L4.5 13.447l1.887 22.597m35.282.4v5.674m0 0H6.402m-.015-6.074v6.022m0-6.022h35.138M33.24 21.706L24 5.882"/>`),
		g.Group(children),
	)
}