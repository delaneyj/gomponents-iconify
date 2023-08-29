package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 20h-2V9H12V7h11a2 2 0 0 1 2 2Z"/><path fill="currentColor" d="M9 23V2H7v5H2v2h5v14a2 2 0 0 0 2 2h14v5h2v-5h5v-2Z"/>`),
		g.Group(children),
	)
}