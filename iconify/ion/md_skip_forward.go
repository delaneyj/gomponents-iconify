package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdSkipForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64 64v384l277.3-192L64 64z" fill="currentColor"/><path d="M384 64h64v384h-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}