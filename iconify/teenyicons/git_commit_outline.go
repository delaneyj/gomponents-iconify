package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommitOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 10.5a3 3 0 0 1 0-6m0 6a3 3 0 0 0 0-6m0 6V15m0-10.5V0"/>`),
		g.Group(children),
	)
}