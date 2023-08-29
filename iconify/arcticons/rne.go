package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 22.974a3.159 3.159 0 0 1 3.159-3.159h0m-3.159 0v8.37m21.586-1.594a3.157 3.157 0 0 1-2.744 1.594h0a3.159 3.159 0 0 1-3.159-3.158v-2.053a3.159 3.159 0 0 1 3.159-3.159h0A3.159 3.159 0 0 1 35 22.974V24h-6.317m-3.024 4.185v-5.211a3.159 3.159 0 0 0-3.159-3.159h0a3.159 3.159 0 0 0-3.159 3.159v5.211m.001-5.211v-3.159"/>`),
		g.Group(children),
	)
}