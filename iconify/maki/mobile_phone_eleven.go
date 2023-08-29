package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilePhoneEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7 2v-.5a.5.5 0 0 0-1 0V2H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zM5 9H4V8h1zm0-2H4V6h1zm2 2H6V8h1zm0-2H6V6h1zm0-2H4V3h3z" fill="currentColor"/>`),
		g.Group(children),
	)
}