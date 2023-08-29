package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30.336 12.547l-10.172-1.074L16 2.133l-4.164 9.34l-10.172 1.074l7.598 6.848L7.14 29.398L16 24.29l8.86 5.11l-2.122-10.004z"/>`),
		g.Group(children),
	)
}