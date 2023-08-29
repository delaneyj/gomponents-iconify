package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HackerNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h24v24H0zm12.8 13.446l4.339-8.303h-1.871q-2.143 4.018-2.839 5.786l-.375.96l-.32-.75a49.079 49.079 0 0 0-3.022-6.243l.129.243H6.857l4.286 8.2v5.52H12.8z"/>`),
		g.Group(children),
	)
}