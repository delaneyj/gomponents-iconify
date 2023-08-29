package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M23.777 15.479A8.64 8.64 0 0 0 16 10a8.64 8.64 0 0 0-7.777 5.479L8 16l.223.521A8.64 8.64 0 0 0 16 22a8.64 8.64 0 0 0 7.777-5.479L24 16ZM16 20a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/><path fill="currentColor" d="m28.504 8.136l-12-7a1 1 0 0 0-1.008 0l-12 7A1 1 0 0 0 3 9v14a1 1 0 0 0 .496.864l12 7a1 1 0 0 0 1.008 0l12-7A1 1 0 0 0 29 23V9a1 1 0 0 0-.496-.864ZM27 22.426l-11 6.416l-11-6.416V9.574l11-6.416l11 6.416Z"/>`),
		g.Group(children),
	)
}