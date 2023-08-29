package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsFileExplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.744 31.55c.614 1.038 1.545 1.425 2.936 1.425h1.925c1.79 0 3.499-1.452 3.814-3.243l.003-.014c.316-1.791-.88-3.243-2.671-3.243h-2.123c-1.793 0-2.99-1.454-2.675-3.247h0c.317-1.797 2.03-3.253 3.828-3.253h1.914c1.39 0 2.322.386 2.936 1.424m-18.116 5.076h4.238m1.116 6.5h-6.5l2.292-13h6.5"/>`),
		g.Group(children),
	)
}