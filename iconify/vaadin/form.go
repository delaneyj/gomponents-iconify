package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Form(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 2v2H6V2h9zm1-1H5v4h11V1zM0 1h4v4H0V1zm15 6v2H6V7h9zm1-1H5v4h11V6zM0 6h4v4H0V6zm15 6v2H6v-2h9zm1-1H5v4h11v-4zM0 11h4v4H0v-4z"/>`),
		g.Group(children),
	)
}