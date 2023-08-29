package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-5L9 4H4Zm7 5a1 1 0 1 0-2 0v1.586l-.293-.293a1 1 0 1 0-1.414 1.414l2 2l.002.002a.997.997 0 0 0 1.41 0l.002-.002l2-2a1 1 0 0 0-1.414-1.414l-.293.293V9Z"/>`),
		g.Group(children),
	)
}