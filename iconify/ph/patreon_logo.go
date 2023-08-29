package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 40a64 64 0 1 0 64 64a64.07 64.07 0 0 0-64-64Zm0 112a48 48 0 1 1 48-48a48.05 48.05 0 0 1-48 48ZM80 40H64a16 16 0 0 0-16 16v152a16 16 0 0 0 16 16h16a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm0 168H64V56h16v152Z"/>`),
		g.Group(children),
	)
}