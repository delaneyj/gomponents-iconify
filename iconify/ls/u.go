package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func U(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M394 654v-47c-42 40-110 64-172 64s-108-24-150-64C28 564 0 505 0 439V128h72v315c2 87 63 157 150 157c88 0 169-70 172-157V128h72v526h-72z"/>`),
		g.Group(children),
	)
}