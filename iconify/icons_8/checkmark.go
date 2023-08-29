package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.28 6.28L11 23.563l-7.28-7.28l-1.44 1.437l8 8l.72.686l.72-.687l18-18l-1.44-1.44z"/>`),
		g.Group(children),
	)
}