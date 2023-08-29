package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m9.5 5l-.313.406L6 9.656V27h19V5zm1 2H23v18H8V10.344zM13 9v4h2V9zm3 0v4h2V9zm3 0v4h2V9zm-9 1v4h2v-4z"/>`),
		g.Group(children),
	)
}