package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v10.719h10.719V0zm13.279 0v10.719h10.719V0zM0 13.281V24h10.719V13.281zm13.279 0V24h10.719V13.281z"/>`),
		g.Group(children),
	)
}