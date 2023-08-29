package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18 5l-6 22h2l6-22zM7.937 6.406l-6.75 9L.75 16l.438.594l6.75 9l1.625-1.188L3.25 16l6.313-8.406zm16.125 0l-1.625 1.188L28.75 16l-6.313 8.406l1.625 1.188l6.75-9L31.25 16l-.438-.594z"/>`),
		g.Group(children),
	)
}