package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 10H7c0-2 1.2-2.6 2-3c.3-.1.5-.2.7-.4c.1-.1.3-.3.1-.7c-.2-.5-.8-1-1.7-1c-1.4 0-1.6 1.2-1.7 1.5l-2-.3C4.5 5 5.4 2.9 8 2.9c1.6 0 3 .9 3.6 2.2c.4 1.1.2 2.2-.6 3c-.4.4-.8.6-1.2.7c-.6.4-.8.2-.8 1.2z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7zm0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8z"/><path fill="currentColor" d="M6.9 11h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}