package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 18a1 1 0 0 1-1-1v-7H7a1 1 0 0 1-.707-1.707l5-5a1 1 0 0 1 1.414 0l5 5A1 1 0 0 1 17 10h-4v7a1 1 0 0 1-1 1Zm-6 1a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}