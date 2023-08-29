package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFooterRemoveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M23 6.5a5.5 5.5 0 1 0-11 0a5.5 5.5 0 0 0 11 0zm-7.976-1.769a.5.5 0 0 1 .707-.707l1.77 1.77l1.771-1.77a.5.5 0 0 1 .713.005a.5.5 0 0 1-.006.702l-1.77 1.77l1.768 1.769a.5.5 0 0 1 .018.688l-.005.005a.5.5 0 0 1-.72.014L17.5 7.208l-1.766 1.767a.5.5 0 1 1-.707-.708l1.766-1.766l-1.769-1.77z" fill="currentColor"/><path d="M8.5 15.75a.25.25 0 0 1 .25-.25h6.5a.25.25 0 0 1 .25.25v1.5a.25.25 0 0 1-.25.25h-6.5a.25.25 0 0 1-.25-.25v-1.5z" fill="currentColor"/><path d="M17.5 13a6.48 6.48 0 0 0 2.495-.496v7.24a2.25 2.25 0 0 1-2.096 2.245l-.154.005h-11.5A2.25 2.25 0 0 1 4 19.898l-.005-.154V4.246A2.25 2.25 0 0 1 6.091 2l.154-.005h6.569A6.5 6.5 0 0 0 17.5 13zm-8.75 1A1.75 1.75 0 0 0 7 15.75v1.5c0 .967.784 1.75 1.75 1.75h6.5A1.75 1.75 0 0 0 17 17.25v-1.5A1.75 1.75 0 0 0 15.25 14h-6.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}