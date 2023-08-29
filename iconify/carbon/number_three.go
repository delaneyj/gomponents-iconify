package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 9h-6v2h6v4h-4v2h4v4h-6v2h6a2 2 0 0 0 2-2V11a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}