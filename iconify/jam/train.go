package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15h8a2 2 0 0 0 2-2v-2H2v2a2 2 0 0 0 2 2zm8.813 1.917A2 2 0 0 1 11.131 20H4.87a2 2 0 0 1-1.682-3.083A4.001 4.001 0 0 1 0 13V4a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v9a4.001 4.001 0 0 1-3.187 3.917zM14 9V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v5h12zm-2-6a1 1 0 0 1 1 1v3a1 1 0 0 1-2 0V4a1 1 0 0 1 1-1zM4 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm8 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm-7.131 4h6.262l-.666-1h-4.93l-.666 1z"/>`),
		g.Group(children),
	)
}