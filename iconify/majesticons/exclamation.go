package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M9.378 4.661c1.143-2.057 4.101-2.057 5.245 0l6.6 11.882c1.111 2-.335 4.457-2.622 4.457H5.399c-2.287 0-3.733-2.457-2.622-4.457l6.6-11.882zM12 8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zm1 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}