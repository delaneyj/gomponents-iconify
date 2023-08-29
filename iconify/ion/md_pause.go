package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 448h106.7V64H96v384zM309.3 64v384H416V64H309.3z" fill="currentColor"/>`),
		g.Group(children),
	)
}