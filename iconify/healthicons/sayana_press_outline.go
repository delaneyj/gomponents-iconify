package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SayanaPressOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m25 4l-2 2v5h-1a1 1 0 0 0-1 1v5h-2v2h-3a4 4 0 0 0-4 4v17a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V23a4 4 0 0 0-4-4h-3v-2h-2v-5a1 1 0 0 0-1-1h-1V4Zm0 7h-2v8h2v-8ZM14 23a2 2 0 0 1 2-2h6a1 1 0 0 0 1 1v3.083a6.002 6.002 0 0 0 0 11.834V39h2v-2.083a6.002 6.002 0 0 0 0-11.834V22a1 1 0 0 0 1-1h6a2 2 0 0 1 2 2v17a2 2 0 0 1-2 2H16a2 2 0 0 1-2-2V23Zm14 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}