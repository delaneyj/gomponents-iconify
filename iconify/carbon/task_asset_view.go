package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskAssetView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.5 26a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5Zm0-5a1.5 1.5 0 1 0 1.5 1.5a1.502 1.502 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M22.5 31a8.5 8.5 0 1 1 8.5-8.5a8.51 8.51 0 0 1-8.5 8.5Zm0-15a6.5 6.5 0 1 0 6.5 6.5a6.507 6.507 0 0 0-6.5-6.5Z"/><path fill="currentColor" d="M25 5h-3V4a2.006 2.006 0 0 0-2-2h-8a2.006 2.006 0 0 0-2 2v1H7a2.006 2.006 0 0 0-2 2v21a2.006 2.006 0 0 0 2 2h5v-2H7V7h3v3h12V7h3v5h2V7a2.006 2.006 0 0 0-2-2Zm-5 3h-8V4h8Z"/>`),
		g.Group(children),
	)
}