package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResidentialCommunityEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.8 8c.7-.1 1.2-.8 1.2-1.5C10 5.7 9.3 5 8.5 5S7 5.7 7 6.5c0 .7.5 1.4 1.2 1.5v1.5H6V1H2v8.5H1v.5h9v-.5H8.8V8zM3 2h2v1H3V2zm0 2h2v1H3V4zm0 2h2v1H3V6z" fill="currentColor"/>`),
		g.Group(children),
	)
}