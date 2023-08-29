package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.906 4.78l-10.5 10.5l-.718.72l.718.72l10.5 10.5l1.406-1.44L7.533 16l9.78-9.78l-1.406-1.44zm7 0l-10.5 10.5l-.72.72l.72.72l10.5 10.5l1.407-1.44L14.53 16l9.783-9.78l-1.407-1.44z"/>`),
		g.Group(children),
	)
}