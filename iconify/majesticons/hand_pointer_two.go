package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 22C4.7 22 3 16.333 3 13.5V10c0-.333.2-1 1-1s1 .667 1 1v2c0 .5.3 1.5 1.5 1.5S8 12.5 8 12V3c0-.333.2-1 1-1s1 .667 1 1v7c.5.5.8 1.2 2 0V8c0-.333.2-1 1-1s1 .667 1 1v1h1c0-.333.2-1 1-1s1 .667 1 1v1h1c0-.333.2-1 1-1s1 .667 1 1v3.5c0 2.833-1.7 8.5-8.5 8.5z"/>`),
		g.Group(children),
	)
}