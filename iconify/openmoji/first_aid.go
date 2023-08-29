package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ea5a47" d="M27.67 43.857H13.72v-15.39l13.994.035l.018-14.118h15.729v13.95l13.817-.082v15.605H43.461v14.084H27.635l.035-14.084z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M27.67 43.857H13.72v-15.39l13.994.035l.018-14.118h15.729v13.95l13.817-.082v15.605H43.461v14.084H27.635l.035-14.084z"/>`),
		g.Group(children),
	)
}