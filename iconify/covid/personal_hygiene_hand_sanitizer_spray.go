package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandSanitizerSpray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.456 8.843H4.281a2.058 2.058 0 0 0-2.058 2.058v10.291c0 1.137.921 2.058 2.058 2.058h6.175a2.058 2.058 0 0 0 2.058-2.058V10.901a2.058 2.058 0 0 0-2.058-2.058ZM5.079 1.871h4.58a.8.8 0 0 1 .8.8v6.172H4.282V2.668a.8.8 0 0 1 .8-.8l-.003.003Zm2.29 12.117v4.116m-2.058-2.058h4.116"/><path d="M21.356 5.71a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m-1.217 5.052a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75m0-7.67a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75M16.21 4.118a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m0 4.118a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}