package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmpStoriesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20V4h10v16H7Zm-4-2V6h2v12H3Zm16 0V6h2v12h-2ZM9 18h6V6H9v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}