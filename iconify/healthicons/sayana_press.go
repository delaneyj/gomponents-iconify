package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SayanaPress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m25 4l-2 2v5h-2a1 1 0 0 0-1 1v5h-1v2h-3a4 4 0 0 0-4 4v17a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V23a4 4 0 0 0-4-4h-3v-2h-1v-5a1 1 0 0 0-1-1h-2V4Zm0 7h-2v8h-1v1a1 1 0 0 0 1 1v3h2v-3a1 1 0 0 0 1-1v-1h-1v-8Zm-1 25a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm-1 2v2h2v-2h-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}