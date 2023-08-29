package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mkr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#1AAB9B"/><path fill="#FFF" d="M7.558 21.156v-8.045l6.101 4.592v3.453h1.558V17.38a.909.909 0 0 0-.363-.726l-7.399-5.569A.909.909 0 0 0 6 11.81v9.346h1.558zm16.874 0v-8.045l-6.101 4.592v3.453h-1.558V17.38c0-.286.134-.555.362-.726l7.4-5.569a.909.909 0 0 1 1.455.726v9.346h-1.558z"/></g>`),
		g.Group(children),
	)
}