package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kickstarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.168 0c-3.2 0-5.797 2.579-5.797 5.758v12.484C1.371 21.42 3.968 24 7.168 24a5.797 5.797 0 0 0 4.768-2.479l.794.79c2.26 2.245 5.943 2.245 8.203 0a5.724 5.724 0 0 0 1.696-4.075a5.724 5.724 0 0 0-1.696-4.074l-2.182-2.168l2.182-2.156a5.724 5.724 0 0 0 1.696-4.074a5.724 5.724 0 0 0-1.696-4.074c-2.26-2.246-5.942-2.246-8.203 0l-.794.789A5.797 5.797 0 0 0 7.168 0Z"/>`),
		g.Group(children),
	)
}