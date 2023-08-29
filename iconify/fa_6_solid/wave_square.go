package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 64c0-17.7 14.3-32 32-32h160c17.7 0 32 14.3 32 32v352h96V256c0-17.7 14.3-32 32-32h128c17.7 0 32 14.3 32 32s-14.3 32-32 32h-96v160c0 17.7-14.3 32-32 32H320c-17.7 0-32-14.3-32-32V96h-96v160c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32h96V64z"/>`),
		g.Group(children),
	)
}