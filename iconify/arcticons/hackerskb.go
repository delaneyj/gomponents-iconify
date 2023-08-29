package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hackerskb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.45 5.5a2 2 0 0 0-1.95 2v33.1a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95Zm5.14 24.62h5.94m-5.94-11.88h5.94m-5.94 5.94h3.87m-3.87-5.94v11.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.41 27.73a3 3 0 0 1-2.94 2.57h0a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h0a3 3 0 0 1 2.94 2.7m-8.51-2.4h-3.78a1.88 1.88 0 0 0-1.88 1.88h0a1.88 1.88 0 0 0 1.88 1.88H25a1.88 1.88 0 0 1 1.88 1.88h0A1.88 1.88 0 0 1 25 30.12h-3.76"/>`),
		g.Group(children),
	)
}