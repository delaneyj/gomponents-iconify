package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 2.133l-4.164 9.34l-10.172 1.074l7.598 6.848L7.14 29.398L16 24.29z"/>`),
		g.Group(children),
	)
}