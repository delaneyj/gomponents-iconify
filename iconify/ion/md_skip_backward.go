package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdSkipBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M170.7 256L448 448V64L170.7 256z" fill="currentColor"/><path d="M64 64h64v384H64z" fill="currentColor"/>`),
		g.Group(children),
	)
}