package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailWarningSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.268 5H3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1h-1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8.876l3.079 1.657c-.1-.356-.103-.745.017-1.126L2 7.74V7a1 1 0 0 1 1-1h3.769l.5-1ZM6.107 9.552l3.496-6.998a1 1 0 0 1 1.79 0l3.5 6.998A1 1 0 0 1 13.998 11H7.002a1 1 0 0 1-.895-1.448Zm4.39-5.557a.5.5 0 0 0-.5.5v3.002a.5.5 0 0 0 1.001 0V4.495a.5.5 0 0 0-.5-.5Zm.501 5.504a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Z"/>`),
		g.Group(children),
	)
}