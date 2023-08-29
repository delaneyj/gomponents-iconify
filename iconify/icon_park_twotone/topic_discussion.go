package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopicDiscussion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTopicDiscussion0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M44 8H4v30h15l5 5l5-5h15V8Z"/><path d="m21 15l-1 17m8-17l-1 17m6-12H16m16 7H15"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTopicDiscussion0)"/>`),
		g.Group(children),
	)
}