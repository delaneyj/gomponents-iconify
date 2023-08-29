package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Txt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 11h3v12h2V11h3V9h-8v2zm-1-2h-2l-2 6l-2-6h-2l2.75 7L12 23h2l2-6l2 6h2l-2.75-7L20 9zM3 11h3v12h2V11h3V9H3v2z"/>`),
		g.Group(children),
	)
}