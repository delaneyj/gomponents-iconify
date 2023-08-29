package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GmailSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.712 3.97a59.38 59.38 0 0 1 10.576 0l1.518.136A3.28 3.28 0 0 1 21.76 6.9a35.257 35.257 0 0 1 0 10.2a3.28 3.28 0 0 1-2.954 2.793l-1.518.136a59.38 59.38 0 0 1-10.576 0l-1.518-.136A3.28 3.28 0 0 1 2.24 17.1a35.257 35.257 0 0 1 0-10.2a3.28 3.28 0 0 1 2.954-2.794l1.518-.135Zm-.856 2.87a.75.75 0 0 0-1.106.66V17a.75.75 0 0 0 1.5 0V8.756l5.394 2.904c.222.12.49.12.712 0l5.394-2.904V17a.75.75 0 0 0 1.5 0V7.5a.75.75 0 0 0-1.106-.66L12 10.148L5.856 6.84Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}