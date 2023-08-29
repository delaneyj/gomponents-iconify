package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 370.43L96 279v98.42l160 88.88l160-88.88V279l-160 91.43z"/><path fill="currentColor" d="M512.25 192L256 45.57L-.25 192L256 338.43l208-118.86V384h48V192.14l.25-.14z"/>`),
		g.Group(children),
	)
}