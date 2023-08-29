package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="4.7" height="6.25" x="23.01" y="20.35" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.71 22.7v3.9m-7.62-.69a2.31 2.31 0 0 1-1.66.69h0a2.35 2.35 0 0 1-2.35-2.35V22.7a2.35 2.35 0 0 1 2.35-2.35h0a2.35 2.35 0 0 1 2.35 2.35v.71m0 .07h-4.7"/><rect width="4.7" height="6.11" x="34.15" y="20.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.15 22.84a2.35 2.35 0 0 1 2.35-2.35h0a2.35 2.35 0 0 1 2.35 2.35v3.76m-4.7-3.76v3.76m22.06-9.39v9.39m-1.23-6.11h1.23m0 0h1.22"/>`),
		g.Group(children),
	)
}