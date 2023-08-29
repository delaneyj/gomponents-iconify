package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM4 17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-11v15m-8-5.027C6.194 19.324 8.86 21 12 21c3.14 0 5.807-1.676 8-5.027M16 10H8"/>`),
		g.Group(children),
	)
}