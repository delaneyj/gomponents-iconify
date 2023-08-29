package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22.775 8A9 9 0 0 1 23 10h-9V1a9 9 0 0 1 8.775 7Zm-2.067 0A6.999 6.999 0 0 0 16 3.292V8h4.708Z"/><path d="M1 14a9 9 0 0 1 11-8.775V12h6.775A9 9 0 1 1 1 14Zm15.803 0H10V7.196A6.804 6.804 0 1 0 16.803 14Z"/></g>`),
		g.Group(children),
	)
}