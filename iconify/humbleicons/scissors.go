package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6c-3.573 2.225-5.943 3.854-8.55 6M20 18c-2.626-1.636-4.602-2.949-6.5-4.382M8.598 9.54a3 3 0 1 0-3.196-5.08a3 3 0 0 0 3.196 5.08Zm0 0A89.3 89.3 0 0 0 11.45 12m-2.852 2.46a3 3 0 1 0-3.196 5.079a3 3 0 0 0 3.196-5.078Zm0 0A89.287 89.287 0 0 1 11.45 12"/>`),
		g.Group(children),
	)
}