package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="22" cy="24" r="2" fill="currentColor"/><path fill="currentColor" d="M29.777 23.479A8.64 8.64 0 0 0 22 18a8.64 8.64 0 0 0-7.777 5.479L14 24l.223.521A8.64 8.64 0 0 0 22 30a8.64 8.64 0 0 0 7.777-5.479L30 24ZM22 28a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><circle cx="8" cy="8" r="1" fill="currentColor"/><circle cx="8" cy="16" r="1" fill="currentColor"/><circle cx="8" cy="24" r="1" fill="currentColor"/><path fill="currentColor" d="M5 21h7v-2H5v-6h16v3h2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h7v-2H5ZM5 5h16v6H5Z"/>`),
		g.Group(children),
	)
}