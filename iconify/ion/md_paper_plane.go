package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdPaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M48 270.9l118.9 44.6L181.7 464 256 360l104 104L464 48 48 270.9zm294.9 126L260 313.4 374.9 152 193.6 289.8 124.9 265l291-156.2-73 288.1z" fill="currentColor"/>`),
		g.Group(children),
	)
}