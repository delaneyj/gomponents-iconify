package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414L5.68 7.094a7.952 7.952 0 0 0-1.248 2.308A5.501 5.501 0 0 0 6.5 20h12.085l.708.707a1 1 0 0 0 1.414-1.414l-1.023-1.023a.978.978 0 0 0-.015-.015L7.757 6.343l-3.05-3.05Zm11.536 4.464a5.992 5.992 0 0 0-5.302-1.663a1 1 0 1 1-.35-1.97a7.992 7.992 0 0 1 7.066 2.22a7.971 7.971 0 0 1 2.307 4.9a4.501 4.501 0 0 1 2.93 5.233a1 1 0 0 1-1.953-.433a2.5 2.5 0 0 0-2.082-3.019a1 1 0 0 1-.86-.995a5.978 5.978 0 0 0-1.756-4.273Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}