package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 44a60 60 0 1 0 60 60a60.07 60.07 0 0 0-60-60Zm0 112a52 52 0 1 1 52-52a52.06 52.06 0 0 1-52 52ZM80 44H64a12 12 0 0 0-12 12v152a12 12 0 0 0 12 12h16a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 164a4 4 0 0 1-4 4H64a4 4 0 0 1-4-4V56a4 4 0 0 1 4-4h16a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}