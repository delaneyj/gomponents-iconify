package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.5 24a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0ZM32 20.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM8.25 8A4.25 4.25 0 0 0 4 12.25v23.5A4.25 4.25 0 0 0 8.25 40h31.5A4.25 4.25 0 0 0 44 35.75v-23.5A4.25 4.25 0 0 0 39.75 8H8.25Zm12.624 12.5A6 6 0 1 1 16 18h16a6 6 0 1 1-4.874 2.5h-6.252Z"/>`),
		g.Group(children),
	)
}