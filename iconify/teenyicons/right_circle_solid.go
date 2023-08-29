package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0ZM6 3.293L10.207 7.5L6 11.707V3.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}