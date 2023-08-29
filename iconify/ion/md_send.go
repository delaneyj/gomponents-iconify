package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M48 448l416-192L48 64v149.333L346 256 48 298.667z" fill="currentColor"/>`),
		g.Group(children),
	)
}