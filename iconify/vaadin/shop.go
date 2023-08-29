package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 15h16v1H0v-1zM0 0v6c.005.732.401 1.37.991 1.715L1 14h9V9h3v5h2V7.72c.599-.35.995-.988 1-1.719V0H0zm4 2h2v4a1 1 0 0 1-2 0V2zM2 7a1 1 0 0 1-1-1V2h2v4a1 1 0 0 1-1 1zm6 5H3V9h5v3zm1-6a1 1 0 0 1-2 0V2h2v4zm3 0a1 1 0 0 1-2 0V2h2v4zm3 0a1 1 0 0 1-2 0V2h2v4z"/>`),
		g.Group(children),
	)
}