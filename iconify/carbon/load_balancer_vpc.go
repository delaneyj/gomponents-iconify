package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadBalancerVpc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 30H2v-6h6zm-4-2h2v-2H4zm15 2h-6v-6h6zm-4-2h2v-2h-2zm15 2h-6v-6h6zm-4-2h2v-2h-2zM16 10a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm10 12h-5a4.005 4.005 0 0 1-4-4h-2a4.005 4.005 0 0 1-4 4H6a2.002 2.002 0 0 0-2 2v4h2v-4h5a5.97 5.97 0 0 0 4-1.54V22h2v-5.54A5.97 5.97 0 0 0 21 18h5v4h2v-4a2.002 2.002 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}