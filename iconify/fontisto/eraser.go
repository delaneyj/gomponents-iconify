package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M25.199 21.599h-9.103l9.952-9.952a1.196 1.196 0 0 0 0-1.696l-9.6-9.6a1.196 1.196 0 0 0-1.696 0L.351 14.75a1.196 1.196 0 0 0 0 1.696l7.2 7.203c.217.217.517.351.848.351h16.834a1.2 1.2 0 1 0-.035-2.399h.002zM15.6 2.896l7.903 7.903l-6.943 6.943l-7.903-7.903z"/>`),
		g.Group(children),
	)
}