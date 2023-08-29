package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17 6v3M7 6v3m5-1c0 2.5.5 4.5.5 4.5l-2 1m-3.5 3c1.444.984 3 1.5 5 1.5s3.556-.516 5-1.5M7 1l-5.7.3L1 7m16-6l5.7.3L23 7m-6 16l5.7-.3l.3-5.7M7 23l-5.7-.3L1 17"/>`),
		g.Group(children),
	)
}