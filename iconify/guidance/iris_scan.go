package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IrisScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m7 1l-5.7.3L1 7m16-6l5.7.3L23 7m-6 16l5.7-.3l.3-5.7M7 23l-5.7-.3L1 17m-1-5c1.276 0 2.5 1 3.5 2l2.606 2.947a8.079 8.079 0 0 0 11.788 0L20.5 14c1-1 2.224-2 3.5-2c-1.276 0-2.5-1-3.5-2l-2.606-2.947a8.079 8.079 0 0 0-11.788 0L3.5 10c-1 1-2.224 2-3.5 2Zm12 2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}