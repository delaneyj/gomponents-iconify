package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ettulo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.5 41.5v-35c0-1.1-.9-2-2-2h-19c-1.1 0-2 .9-2 2v35c0 1.1.9 2 2 2h19c1.1 0 2-.9 2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="15.428" r="2.8"/><path d="M29.4 23.628a5.4 5.4 0 0 0-10.8 0h10.8Z"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.839 38.275V36.56a3.161 3.161 0 1 1 6.322 0v1.714"/><path d="M22.986 37.539v-.793a1.014 1.014 0 0 1 2.027 0v2.597m-2.027-.19v.807"/></g>`),
		g.Group(children),
	)
}