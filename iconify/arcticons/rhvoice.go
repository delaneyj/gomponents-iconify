package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rhvoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2.006 2.006 0 0 1-2-2v-33a2.006 2.006 0 0 1 2-2h33a2.006 2.006 0 0 1 2 2v33a2.006 2.006 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 27.54v-6.703h2.195a2.252 2.252 0 0 1 0 4.503H8m2.195 0l2.194 2.199m1.678-6.702v6.704m4.442-6.704v6.704m-4.442-3.365h4.442m6.036-3.34l-2.221 6.705l-2.221-6.705"/><rect width="3.352" height="4.441" x="25.109" y="23.099" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.676"/><circle cx="30.209" cy="21.046" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.209 23.099v4.442m9.571-.846a1.675 1.675 0 0 1-1.456.846h0a1.676 1.676 0 0 1-1.676-1.676v-1.09a1.676 1.676 0 0 1 1.676-1.676h0A1.676 1.676 0 0 1 40 24.775v.545h-3.352m-1.574 1.376a1.675 1.675 0 0 1-1.455.845h0a1.676 1.676 0 0 1-1.676-1.676v-1.09a1.676 1.676 0 0 1 1.676-1.676h0a1.675 1.675 0 0 1 1.454.841"/>`),
		g.Group(children),
	)
}