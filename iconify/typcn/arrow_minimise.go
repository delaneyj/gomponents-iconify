package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMinimise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.121 13a1 1 0 1 0 0 2h1.465l-3.293 3.293a.999.999 0 1 0 1.414 1.414l3.414-3.414V18c0 .552.447 1 1 1s.879-.448.879-1v-5H6.121zM7 11a1 1 0 0 0 1-1V8h2a1 1 0 1 0 0-2H6.001L6 10a1 1 0 0 0 1 1zm10 2a1 1 0 0 0-1 1v2h-2a1 1 0 1 0 0 2h4v-4a1 1 0 0 0-1-1zm1.293-8.707L15 7.586V6a1 1 0 1 0-2 0v5h5a1 1 0 0 0 0-2h-1.586l3.293-3.292c.391-.391.391-1.023 0-1.414s-1.023-.392-1.414-.001z"/>`),
		g.Group(children),
	)
}