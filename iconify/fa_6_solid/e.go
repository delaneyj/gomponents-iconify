package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func E(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h224c17.7 0 32-14.3 32-32s-14.3-32-32-32H64V288h160c17.7 0 32-14.3 32-32s-14.3-32-32-32H64V96h224c17.7 0 32-14.3 32-32s-14.3-32-32-32H64z"/>`),
		g.Group(children),
	)
}