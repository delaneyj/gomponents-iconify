package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 4.703c0-1.857-2.31-2.711-3.519-1.301L5.84 7.65a1 1 0 0 1-.76.35H5a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h.08a1 1 0 0 1 .76.35l3.641 4.248c1.209 1.41 3.519.556 3.519-1.301V4.703zm-5.642 4.25L11 4.702v14.594l-3.642-4.25A3 3 0 0 0 5.08 14H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h.08a3 3 0 0 0 2.278-1.048zm10.35.34a1 1 0 1 0-1.415 1.414L17.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L19 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L20.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L19 10.586l-1.293-1.293z"/></g>`),
		g.Group(children),
	)
}