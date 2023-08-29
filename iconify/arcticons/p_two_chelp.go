package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PTwoChelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.532 8.094l-2.17 3.004a.99.99 0 0 1-.804.414c-1.726 0-7.432.003-10.268.002a1.79 1.79 0 0 0-1.79 1.792V38.53a1.79 1.79 0 0 0 1.79 1.791h35.418a1.79 1.79 0 0 0 1.791-1.793V13.305a1.79 1.79 0 0 0-1.79-1.793h-10.26a1 1 0 0 1-.812-.415L28.47 8.095a1 1 0 0 0-.81-.415h-7.316a1 1 0 0 0-.81.414Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.765 29.42l-4.339-6.244l-6.42 9.238H36.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.93 32.414l-9.684-13.936l-6.542 9.414"/>`),
		g.Group(children),
	)
}