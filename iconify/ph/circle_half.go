package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm8 16.37a86.4 86.4 0 0 1 16 3v169.3a86.4 86.4 0 0 1-16 3Zm32 9.26a87.81 87.81 0 0 1 16 10.54v135.66a87.81 87.81 0 0 1-16 10.54ZM40 128a88.11 88.11 0 0 1 80-87.63v175.26A88.11 88.11 0 0 1 40 128Zm160 50.54V77.46a87.82 87.82 0 0 1 0 101.08Z"/>`),
		g.Group(children),
	)
}