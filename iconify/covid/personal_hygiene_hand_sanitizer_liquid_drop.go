package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerLiquidDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.251 21.664h3a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v6m6-8.001h-4V5.438a.775.775 0 0 1 .775-.775h2.45a.775.775 0 0 1 .775.775v2.225Zm2 4.001h-8m8 6h-5M.749 18.836l2.937 2.35a5.247 5.247 0 0 0 3.279 1.15H13a1.75 1.75 0 0 0 1.75-1.75v0a1.75 1.75 0 0 0-1.75-1.75H8.478"/><path d="M.908 13.586h3.5c1.037 0 2.051.308 2.914.883l2 1.333a1.633 1.633 0 0 1 .518 2.307v0a1.634 1.634 0 0 1-2.089.555l-2.627-1.578m7.127-7.922a2.5 2.5 0 0 1-5 0c0-1.875 2.5-5 2.5-5s2.5 3.125 2.5 5Zm7-7.5v3m-6-1l.447-.895a2 2 0 0 1 1.789-1.105h6.764"/></g>`),
		g.Group(children),
	)
}