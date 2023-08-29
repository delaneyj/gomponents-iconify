package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8h8v8H8z" opacity=".3"/><path fill="currentColor" d="M6 18h12V6H6v12zM8 8h8v8H8V8z"/>`),
		g.Group(children),
	)
}