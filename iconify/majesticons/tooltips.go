package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tooltips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M12 2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 11.07 12H12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z"/><path d="M18 10h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-.93a2 2 0 0 0-1.664.89l-.574.862a1 1 0 0 1-1.664 0l-.574-.861A2 2 0 0 0 12.93 20H12a2 2 0 0 1-2-2v-1"/></g>`),
		g.Group(children),
	)
}