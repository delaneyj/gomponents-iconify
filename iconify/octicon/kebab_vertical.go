package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KebabVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M0 2.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0zm0 5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0zM1.5 14a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z" fill="currentColor"/>`),
		g.Group(children),
	)
}