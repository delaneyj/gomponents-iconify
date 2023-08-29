package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Display(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1v10h16V1H0zm15 9H1V2h14v8zm-4.5 2h-5L5 14l-1 1h8l-1-1z"/>`),
		g.Group(children),
	)
}