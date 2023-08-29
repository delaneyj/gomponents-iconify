package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommonWorkflowLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16.083 12.271l-4.755 4.76l4.693 4.682l2.063-2.063l-2.635-2.62l2.698-2.698zm2.058 17.625l-2.672-2.729l5.047-5.089l-2.078-2.094l-7.109 7.203l.156.151l-.021.01l4.609 4.646zm2.531-18.083l-4.917-4.724l4.917-4.964L18.583 0l-7.12 7.135l.146.135v.01l7 6.589z"/>`),
		g.Group(children),
	)
}