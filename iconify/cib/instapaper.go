package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instapaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.688 27.01c0 2.427.359 2.786 3.911 3.057v1.932H8.401v-1.932c3.557-.271 3.911-.63 3.911-3.057V4.942c0-2.375-.359-2.786-3.911-3.057V0h15.198v1.885c-3.552.271-3.911.677-3.911 3.057z"/>`),
		g.Group(children),
	)
}