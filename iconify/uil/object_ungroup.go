package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectUngroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18.28v-6.56A2 2 0 1 0 18.28 9H15V5.72A2 2 0 1 0 12.28 3H5.72A2 2 0 1 0 3 5.72v6.56A2 2 0 1 0 5.72 15H9v3.28A2 2 0 1 0 11.72 21h6.56A2 2 0 1 0 21 18.28ZM8 10a2 2 0 0 0 1 1.72V13H5.72a1.91 1.91 0 0 0-.72-.72V5.72A1.91 1.91 0 0 0 5.72 5h6.56a1.91 1.91 0 0 0 .72.72V9h-1.28A2 2 0 0 0 8 10Zm5 1v1.28a1.91 1.91 0 0 0-.72.72H11v-1.28a1.91 1.91 0 0 0 .72-.72Zm6 7.28a1.91 1.91 0 0 0-.72.72h-6.56a1.91 1.91 0 0 0-.72-.72V15h1.28A2 2 0 1 0 15 12.28V11h3.28a1.91 1.91 0 0 0 .72.72Z"/>`),
		g.Group(children),
	)
}