package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 3C13.489 3 9 7.489 9 13c0 2.395.839 4.587 2.25 6.313L3.281 27.28l1.438 1.44l7.968-7.969A9.923 9.923 0 0 0 19 23c5.511 0 10-4.489 10-10S24.511 3 19 3zm0 2c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8zm0 3a4 4 0 0 0-4 4c0 3 4 7 4 7s4-4 4-7a4 4 0 0 0-4-4zm0 2a2 2 0 1 1 .001 3.999A2 2 0 0 1 19 10z"/>`),
		g.Group(children),
	)
}