package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M38.5 24a5 5 0 1 0-10 0a5 5 0 0 0 10 0Zm5.5 0c0-5.523-4.477-10-10-10H14C8.477 14 4 18.477 4 24s4.477 10 10 10h20c5.523 0 10-4.477 10-10Zm-10-7.5a7.5 7.5 0 0 1 0 15H14a7.5 7.5 0 0 1 0-15h20Z"/>`),
		g.Group(children),
	)
}