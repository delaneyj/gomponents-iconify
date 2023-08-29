package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiDeskclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="26.63" r="16.87" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.943 4.5v22.3m0 0l7.754 5.943m2.653 7.21l2.894 3.487m-23.14-3.147l-2.556 3.147M19.019 4.5h9.622"/>`),
		g.Group(children),
	)
}