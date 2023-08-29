package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trademark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 8v2h4v13h2V10h4V8zm12 0v15h2V10.875l4.156 6.656l.844 1.344l.844-1.344L27 10.875V23h2V8h-2.563l-.28.469L22 15.125l-4.156-6.656L17.562 8z"/>`),
		g.Group(children),
	)
}