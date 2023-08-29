package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbSctFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#0065bd" d="M0 0h640v480H0z"/><path stroke="#fff" stroke-width=".6" d="m0 0l5 3M0 3l5-3" transform="scale(128 160)"/>`),
		g.Group(children),
	)
}