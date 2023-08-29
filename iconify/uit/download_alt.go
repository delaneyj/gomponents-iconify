package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.162 13.655a.5.5 0 0 0-.007.707l3.491 3.492a.498.498 0 0 0 .708 0l3.49-3.492a.5.5 0 0 0-.706-.707L12.5 16.293V2.5a.5.5 0 0 0-1 0v13.793l-2.638-2.638a.5.5 0 0 0-.7 0zM18 9h-1.5a.5.5 0 0 0 0 1H18a2.003 2.003 0 0 1 2 2v7a2.003 2.003 0 0 1-2 2H6a2.003 2.003 0 0 1-2-2v-7a2.003 2.003 0 0 1 2-2h2.5a.5.5 0 0 0 0-1H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}