package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadBalancerPool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 15h12v2H10zM8.7 6.285A2.966 2.966 0 0 0 9 5a3 3 0 1 0-3 3a2.96 2.96 0 0 0 1.285-.3L10 10.413V13h2V9.586zM6 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm13-1a3 3 0 1 0-4 2.815V13h2V7.816A2.996 2.996 0 0 0 19 5zm-3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm10-4a3.003 3.003 0 0 0-3 3a2.966 2.966 0 0 0 .3 1.285l-3.3 3.3V13h2v-2.586L24.715 7.7A2.96 2.96 0 0 0 26 8a3 3 0 0 0 0-6zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1zM12 19h-2v2.586L7.285 24.3A2.96 2.96 0 0 0 6 24a3 3 0 1 0 3 3a2.966 2.966 0 0 0-.3-1.285l3.3-3.3zm-6 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm11-3.816V19h-2v5.184a3 3 0 1 0 2 0zM16 28a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm10-4a2.96 2.96 0 0 0-1.285.3L22 21.587V19h-2v3.414l3.3 3.3A2.966 2.966 0 0 0 23 27a3 3 0 1 0 3-3zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1z"/>`),
		g.Group(children),
	)
}