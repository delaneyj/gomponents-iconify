package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveBoxThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m219.58 70.21l-16-32A4 4 0 0 0 200 36H56a4 4 0 0 0-3.58 2.21l-16 32A4 4 0 0 0 36 72v136a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V72a4 4 0 0 0-.42-1.79ZM58.47 44h139.06l12 24H46.47ZM208 212H48a4 4 0 0 1-4-4V76h168v132a4 4 0 0 1-4 4Zm-45.17-62.83a4 4 0 0 1 0 5.66l-32 32a4 4 0 0 1-5.66 0l-32-32a4 4 0 0 1 5.66-5.66L124 174.34V104a4 4 0 0 1 8 0v70.34l25.17-25.17a4 4 0 0 1 5.66 0Z"/>`),
		g.Group(children),
	)
}