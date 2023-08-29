package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 14a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-9 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0ZM4.75 4A2.75 2.75 0 0 0 2 6.75v14.5A2.75 2.75 0 0 0 4.75 24h18.5A2.75 2.75 0 0 0 26 21.25V6.75A2.75 2.75 0 0 0 23.25 4H4.75Zm7.623 8A3.5 3.5 0 1 1 9.5 10.5h9a3.5 3.5 0 1 1-2.873 1.5h-3.254Z"/>`),
		g.Group(children),
	)
}