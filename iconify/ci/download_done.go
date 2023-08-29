package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19H5v-2h14v2Zm-9-4.58l-4-4l1.41-1.41L10 11.59L16.59 5L18 6.42l-8 8Z"/>`),
		g.Group(children),
	)
}