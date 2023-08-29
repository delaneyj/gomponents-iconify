package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 8h-1V2h-2v6h-6V2h-2v6h-1a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h1v6h2v-6h6v6h2v-6h1a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2zM10 22V10h12v12zm15-8h2v4h-2z"/>`),
		g.Group(children),
	)
}