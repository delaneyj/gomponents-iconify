package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nationwide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.74 39v4.5h10.08v-5.45h8.41l-.1-14.79l-8.58-.05V4.5h-7.42v5.2H22v2.82h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.74 43.5H13.67v-7.75c1.94-3.23 4.7-7.28 7.12-10.86h11.42L25 35.69H13.67m20.92-21.24l8.54 8.81m-7.43 7.79v2.73"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.788 35.554C8.593 34.156 4.77 29.412 4.77 23.776c0-6.736 5.46-12.196 12.196-12.196s12.196 5.46 12.196 12.196c0 .375-.017.747-.05 1.114"/>`),
		g.Group(children),
	)
}