package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KebabVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.444 13.832a1 1 0 1 0 1.111-1.663a1 1 0 0 0-1.11 1.662zM8 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-5a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}