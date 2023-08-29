package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownFilledAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 15h10v2H22zm-4 15a.997.997 0 0 1-.71-.297L9.665 22H3a1 1 0 0 1-1-.999V11a1 1 0 0 1 .999-1h6.667l7.623-7.703A1 1 0 0 1 19 3v26a1 1 0 0 1-1 1z"/>`),
		g.Group(children),
	)
}