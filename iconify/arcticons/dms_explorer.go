package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DmsExplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="16.149" ry="16.149"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.529 26.445l-9.983 5.866c-2.036.738-3.4-.798-3.294-2.651V18.34c-.106-1.853 1.258-3.39 3.294-2.65l9.983 5.866c2.94 1.727 2.994 3.13 0 4.889Z"/>`),
		g.Group(children),
	)
}