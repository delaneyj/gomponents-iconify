package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HackerNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 2v460h460V2H2zm250 265v113h-40V267L111 68h47l74 150l76-150h45z"/>`),
		g.Group(children),
	)
}