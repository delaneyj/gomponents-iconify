package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdVolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64 192v128h85.334L256 431.543V80.458L149.334 192H64z" fill="currentColor"/>`),
		g.Group(children),
	)
}