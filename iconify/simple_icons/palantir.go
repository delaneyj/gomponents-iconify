package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palantir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.147 18L12 21.178L3.853 18L2.5 20.343L12 24l9.5-3.657L20.147 18zM12 0a9.5 9.5 0 1 0 0 19a9.5 9.5 0 0 0 0-19zm0 16.078a6.568 6.568 0 1 1 0-13.136a6.568 6.568 0 0 1 0 13.136z"/>`),
		g.Group(children),
	)
}