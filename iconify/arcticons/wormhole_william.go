package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WormholeWilliam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="19.151" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.5" ry="12.886"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.774 30.838c3.638 4.328 16.881 8.134 16.881 8.134"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.18 28.649c-.719 2.137 2.045 2.867 2.045 7.143s-4.901 5.943-7.821 5.943s-3.494-1.381-5.397-1.381a1.808 1.808 0 0 1-1.929-2.073a1.88 1.88 0 0 1 1.245-1.855"/>`),
		g.Group(children),
	)
}