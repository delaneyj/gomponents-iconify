package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadBalancerNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 16h-8v-4h-2v4H7a2.002 2.002 0 0 0-2 2v4h2v-4h8v4h2v-4h8v4h2v-4a2.002 2.002 0 0 0-2-2zm-5-6V2h-8v8h8zm-6-2V4h4v4zm12 16a2.995 2.995 0 0 0-2.816 2h-4.368a2.982 2.982 0 0 0-5.632 0H8.816a3 3 0 1 0 0 2h4.368a2.982 2.982 0 0 0 5.632 0h4.368A2.995 2.995 0 1 0 26 24zM6 28a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1z"/>`),
		g.Group(children),
	)
}