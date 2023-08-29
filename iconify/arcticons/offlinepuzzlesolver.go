package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Offlinepuzzlesolver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.833 5.5v37m12.334-37v37M42.5 17.833h-37m37 12.334h-37m4.417-15h3.5m-3.5-6.047l1.75-.953m0 0v7m25.236 0a1.75 1.75 0 0 0 1.75-1.75h0a1.75 1.75 0 0 0-1.75-1.75h0a1.75 1.75 0 0 0 1.75-1.75h0a1.75 1.75 0 0 0-1.75-1.75m-2.888 6.409a2.981 2.981 0 0 0 2.177.59h.71m-2.887-6.415a2.981 2.981 0 0 1 2.178-.584l.71.001M24.569 36.333a1.75 1.75 0 0 0 1.75-1.75h0a1.75 1.75 0 0 0-1.75-1.75H23.43a1.75 1.75 0 0 0-1.75 1.75h0a1.75 1.75 0 0 0 1.75 1.75h0a1.75 1.75 0 0 0-1.75 1.75h0a1.75 1.75 0 0 0 1.75 1.75h1.138a1.75 1.75 0 0 0 1.75-1.75h0a1.75 1.75 0 0 0-1.75-1.75M13.105 27.5v-7l-3.757 4.702h4.637"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".998" d="m10.219 39.833l3.766-7H9.348"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.681 26.91a2.834 2.834 0 0 0 2.124.59h.253a2.261 2.261 0 0 0 2.26-2.261h0a2.261 2.261 0 0 0-2.26-2.261H21.68V20.5h4.371"/><circle cx="36.333" cy="25.181" r="2.319" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.447 21.354a2.264 2.264 0 0 0-1.95-.854h-.164a2.319 2.319 0 0 0-2.318 2.319v2.362M21.681 10.485a2.319 2.319 0 1 1 4.638 0a2.163 2.163 0 0 1-.68 1.64c-.937.824-3.958 3.042-3.958 3.042h4.638"/>`),
		g.Group(children),
	)
}