package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 116h-28V48a20 20 0 0 0-20-20H96a20 20 0 0 0-20 20v68H48a12 12 0 0 0 0 24h28v68a20 20 0 0 0 20 20h64a20 20 0 0 0 20-20v-68h28a12 12 0 0 0 0-24Zm-52 88h-56V52h56Z"/>`),
		g.Group(children),
	)
}