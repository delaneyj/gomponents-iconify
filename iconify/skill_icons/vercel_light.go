package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VercelLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#F4F2ED" rx="60"/><path fill="#000" d="m128 34l95 164.853H33L128 34Z"/></g>`),
		g.Group(children),
	)
}