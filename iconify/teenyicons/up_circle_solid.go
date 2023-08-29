package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15ZM3.293 9L7.5 4.793L11.707 9H3.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}