package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 7.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0Zm4.5 0a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0ZM3.5 3A1.5 1.5 0 0 0 2 4.5V11a1.5 1.5 0 0 0 1.5 1.5h9A1.5 1.5 0 0 0 14 11V4.5A1.5 1.5 0 0 0 12.5 3h-9Zm4 4.75A1.75 1.75 0 1 1 5.75 6h4.5a1.75 1.75 0 1 1-1.582 1H7.332c.108.227.168.482.168.75Z"/>`),
		g.Group(children),
	)
}