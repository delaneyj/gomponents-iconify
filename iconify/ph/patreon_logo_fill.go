package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 104a64 64 0 1 1-64-64a64.07 64.07 0 0 1 64 64ZM80 40H64a16 16 0 0 0-16 16v152a16 16 0 0 0 16 16h16a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}