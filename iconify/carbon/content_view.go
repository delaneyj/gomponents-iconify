package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="19" r="2" fill="currentColor"/><path fill="currentColor" d="M23.777 18.479A8.64 8.64 0 0 0 16 13a8.64 8.64 0 0 0-7.777 5.479L8 19l.223.521A8.64 8.64 0 0 0 16 25a8.64 8.64 0 0 0 7.777-5.479L24 19ZM16 23a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="currentColor" d="M27 3H5a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2ZM5 5h22v4H5Zm0 22V11h22v16Z"/>`),
		g.Group(children),
	)
}