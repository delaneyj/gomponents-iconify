package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lsposed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.616 18v12h6m12.818 0V18h3.928a4.03 4.03 0 0 1 0 8.06h-3.928M19.12 28.685A3.357 3.357 0 0 0 22.062 30h1.776a2.997 2.997 0 0 0 2.994-3h0a2.997 2.997 0 0 0-2.994-3h-1.963a2.997 2.997 0 0 1-2.993-3h0a2.997 2.997 0 0 1 2.993-3h1.777a3.356 3.356 0 0 1 2.942 1.315"/>`),
		g.Group(children),
	)
}