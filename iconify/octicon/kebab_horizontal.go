package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KebabHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M1.5 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zM13 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}