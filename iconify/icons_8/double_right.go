package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.094 4.78L7.688 6.22l9.78 9.78l-9.78 9.78l1.406 1.44l10.5-10.5l.718-.72l-.718-.72l-10.5-10.5zm7 0l-1.407 1.44L24.47 16l-9.782 9.78l1.406 1.44l10.5-10.5l.718-.72l-.718-.72l-10.5-10.5z"/>`),
		g.Group(children),
	)
}