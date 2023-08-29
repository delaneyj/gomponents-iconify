package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartAddLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Zm1.243-9.243a6 6 0 0 1 .237 8.235l-1.42-1.418c1.33-1.524 1.26-3.914-.233-5.404a4.001 4.001 0 0 0-5.49-.153l-1.335 1.198l-1.336-1.197a4 4 0 0 0-5.686 5.605l8.432 8.446L12 21.485l-8.48-8.492A6 6 0 0 1 12 4.529a5.998 5.998 0 0 1 8.242.228Z"/>`),
		g.Group(children),
	)
}