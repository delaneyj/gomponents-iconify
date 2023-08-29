package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feShield0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShield1" fill="currentColor" fill-rule="nonzero"><path id="feShield2" d="m6 16.764l6 3V4H6v12.764ZM6 2h12a2 2 0 0 1 2 2v14l-8 4l-8-4V4a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}