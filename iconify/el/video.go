package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 145.898v908.203h1200V145.898H0zm147.144 147.218h905.713v613.77H147.144v-613.77zm318.237 106.861v408.839L818.848 603.81L465.381 399.977z"/>`),
		g.Group(children),
	)
}