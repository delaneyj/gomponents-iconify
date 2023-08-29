package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowCloseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm4.688 3.313l-1.407 1.406L14.562 16l-4.343 4.344l1.406 1.406l4.344-4.344l4.312 4.313l1.407-1.407L17.375 16l4.25-4.25l-1.406-1.406l-4.25 4.25z"/>`),
		g.Group(children),
	)
}