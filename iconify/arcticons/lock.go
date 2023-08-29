package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.17 19.26h27.66a1.9 1.9 0 0 1 1.898 1.9V41.6a1.9 1.9 0 0 1-1.899 1.9H10.171a1.9 1.9 0 0 1-1.9-1.9V21.16a1.9 1.9 0 0 1 1.9-1.9Zm3.42 0V14.9a10.401 10.401 0 0 1 20.802 0v4.36"/><circle cx="23.991" cy="31.38" r="4.034" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}