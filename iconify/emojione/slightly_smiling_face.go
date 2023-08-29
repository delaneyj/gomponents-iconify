package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlightlySmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><g fill="#664e27"><circle cx="20.5" cy="26.6" r="5"/><circle cx="43.5" cy="26.6" r="5"/><path d="M44.6 40.3c-8.1 5.7-17.1 5.6-25.2 0c-1-.7-1.8.5-1.2 1.6c2.5 4 7.4 7.7 13.8 7.7s11.3-3.6 13.8-7.7c.6-1.1-.2-2.3-1.2-1.6"/></g>`),
		g.Group(children),
	)
}