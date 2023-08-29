package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommitSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 7.5a3.5 3.5 0 0 1 3-3.465V0h1v4.035a3.5 3.5 0 0 1 0 6.93V15H7v-4.035A3.5 3.5 0 0 1 4 7.5Z"/>`),
		g.Group(children),
	)
}