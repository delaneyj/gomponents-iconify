package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blcn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#2aabe4"/><path fill="#fff" d="M7.9 7h6.2a.9.9 0 0 1 .9.9v6.2a.9.9 0 0 1-.9.9H7.9a.9.9 0 0 1-.9-.9V7.9a.9.9 0 0 1 .9-.9zm10 0h6.2a.9.9 0 0 1 .9.9v6.2a.9.9 0 0 1-.9.9h-6.2a.9.9 0 0 1-.9-.9V7.9a.9.9 0 0 1 .9-.9zm0 10h6.2a.9.9 0 0 1 .9.9v6.2a.9.9 0 0 1-.9.9h-6.2a.9.9 0 0 1-.9-.9v-6.2a.9.9 0 0 1 .9-.9zm-10 0h6.2a.9.9 0 0 1 .9.9v6.2a.9.9 0 0 1-.9.9H7.9a.9.9 0 0 1-.9-.9v-6.2a.9.9 0 0 1 .9-.9z"/></g>`),
		g.Group(children),
	)
}