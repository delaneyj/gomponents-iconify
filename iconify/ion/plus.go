package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M448 224H288V64h-64v160H64v64h160v160h64V288h160z" fill="currentColor"/>`),
		g.Group(children),
	)
}