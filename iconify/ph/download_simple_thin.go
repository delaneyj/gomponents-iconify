package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 152v56a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12v-56a4 4 0 0 1 8 0v56a4 4 0 0 0 4 4h160a4 4 0 0 0 4-4v-56a4 4 0 0 1 8 0Zm-94.83 2.83a4 4 0 0 0 5.66 0l40-40a4 4 0 1 0-5.66-5.66L132 142.34V40a4 4 0 0 0-8 0v102.34l-33.17-33.17a4 4 0 0 0-5.66 5.66Z"/>`),
		g.Group(children),
	)
}