package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gifmaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2ZM25.21 17.03v13.94"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.82 21.7a4.66 4.66 0 0 0-4.91-4.7a4.84 4.84 0 0 0-4.41 5v4.3a4.67 4.67 0 0 0 4.66 4.7h0a4.66 4.66 0 0 0 4.66-4.67h-4.66m13.37-9.3h6.97M29.53 24h4.53m-4.53-6.97v13.94"/>`),
		g.Group(children),
	)
}