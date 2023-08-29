package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.126 5.633a1 1 0 0 1 1.748 0l6.601 11.881A1 1 0 0 1 18.601 19H5.399a1 1 0 0 1-.874-1.486l6.6-11.881zm3.497-.972c-1.143-2.057-4.102-2.057-5.245 0L2.777 16.543C1.666 18.543 3.112 21 5.399 21h13.202c2.287 0 3.733-2.457 2.622-4.457l-6.6-11.882zM12 8a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1zm-1 8a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}