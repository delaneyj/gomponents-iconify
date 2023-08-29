package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottlesContainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 10V8a1 1 0 0 0-1-1h-3v2h2v1.969s2 1.124 2 3.48V23h-4v2h5a1 1 0 0 0 1-1v-9.55A5.93 5.93 0 0 0 23 10zm-8 0V8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v2a5.93 5.93 0 0 0-2 4.449V24a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-9.551A5.93 5.93 0 0 0 15 10zm0 13H9v-8.551c0-2.356 2-3.48 2-3.48v-1.97h2v1.97s2 1.124 2 3.48V23z"/><path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 26H4V4h24v24Z"/>`),
		g.Group(children),
	)
}