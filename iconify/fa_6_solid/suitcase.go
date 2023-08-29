package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 56v40h160V56c0-4.4-3.6-8-8-8H184c-4.4 0-8 3.6-8 8zm-48 40V56c0-30.9 25.1-56 56-56h144c30.9 0 56 25.1 56 56v424H128V96zm-64 0h32v384H64c-35.3 0-64-28.7-64-64V160c0-35.3 28.7-64 64-64zm384 384h-32V96h32c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64z"/>`),
		g.Group(children),
	)
}