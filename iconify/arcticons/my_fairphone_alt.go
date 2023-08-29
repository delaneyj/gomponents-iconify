package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyFairphoneAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.338 15.908v6.975m-10.83-6.975h3.486m-3.486 3.487h2.266m-2.266-3.487v6.975m13.182-.001v-6.974h2.267a2.354 2.354 0 0 1 0 4.708H21.69m2.432-.006l2.186 2.271m-9.296 0l-2.267-6.974l-2.354 6.976m.786-2.356h3.05m19.855 11.564H39.5m-3.418-6.834H39.5m-3.418 3.418h2.221m-2.221-3.418v6.836M8.5 32.09v-6.833h2.222a2.307 2.307 0 1 1 0 4.614H8.5m20.553 2.219v-6.833l4.527 6.835v-6.835m-18.26 0v6.835m4.529-6.835v6.835m-4.529-3.417h4.529m2.372 1.054a2.264 2.264 0 1 0 4.527 0v-2.305a2.297 2.297 0 0 0-2.306-2.307a2.225 2.225 0 0 0-2.22 2.307Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}