package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Horse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.572 12.36c-1.889-2.274-.475-5.31-2.205-8.448a4.144 4.144 0 0 0-1.564-1.63a.975.975 0 0 0 .033-.03c.916-.928 1.412-1.903 1.106-2.178c-.305-.277-1.296.252-2.212 1.18c-.176.178-.33.356-.472.533c-.976-.081-2.003.242-2.895 1.095c0 0-3.909 3.697-4.949 4.697c-1.044.998.367 1.534.367 1.534s4.003-.967 4.833-.645c1.445.558-2.029 2.992-1.918 6.148c.066 2.044 5.214 1.205 8.562.225c1.825-.533 1.67-2.054 1.314-2.48z"/>`),
		g.Group(children),
	)
}