package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDocs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.35 4.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V14.49h-8a2 2 0 0 1-1.95-2v-8Zm19.26 0l9.99 9.99m-23.76 8.48h16.32m-16.32 12.1h16.32m-16.32-6.05h16.32"/>`),
		g.Group(children),
	)
}