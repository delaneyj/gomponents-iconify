package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 0h1v4.03a4 4 0 0 1 0 7.94V16h-1v-4.03a4 4 0 0 1 0-7.94V0ZM8 10.6a2.6 2.6 0 1 0 0-5.2a2.6 2.6 0 0 0 0 5.2Z"/>`),
		g.Group(children),
	)
}