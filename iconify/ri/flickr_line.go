package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17A5 5 0 1 0 6 7a5 5 0 0 0 0 10Zm3-5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm9 5a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm3-5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`),
		g.Group(children),
	)
}