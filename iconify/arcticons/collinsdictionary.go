package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collinsdictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 19.844h3.099v8.265H9.5zm23.525 8.265v-8.265h1.86A3.616 3.616 0 0 1 38.5 23.46v1.033a3.616 3.616 0 0 1-3.616 3.616Zm-11.263-2.772v.034a2.738 2.738 0 0 1-2.737 2.738h0a2.738 2.738 0 0 1-2.738-2.738v-2.79a2.738 2.738 0 0 1 2.738-2.737h0a2.738 2.738 0 0 1 2.737 2.737v.034m2.788 2.804a2.738 2.738 0 1 0 5.476 0v-2.79a2.738 2.738 0 0 0-5.475 0Z"/>`),
		g.Group(children),
	)
}