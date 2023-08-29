package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ratio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/><rect width="5.797" height="7.681" x="32.203" y="20.442" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.899"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 23.34a2.899 2.899 0 0 1 2.899-2.898h0m-2.899 0v7.681"/><circle cx="29.288" cy="16.891" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.288 20.442v7.681m-4.301-10.072v10.072m-1.522-7.681h3.044m-5.87 4.783a2.899 2.899 0 0 1-2.899 2.898h0a2.899 2.899 0 0 1-2.899-2.898V23.34a2.899 2.899 0 0 1 2.899-2.899h0a2.899 2.899 0 0 1 2.899 2.899m0 4.783v-7.681"/>`),
		g.Group(children),
	)
}