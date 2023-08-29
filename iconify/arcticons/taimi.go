package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taimi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2zm-31 14.516h5.63m-2.762 8.499v-8.499"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.804 26.39a2.13 2.13 0 0 1-2.124 2.125a2.13 2.13 0 0 1-2.125-2.125v-1.38a2.13 2.13 0 0 1 2.125-2.125a2.13 2.13 0 0 1 2.124 2.124m0 3.506v-5.63"/><circle cx="22.991" cy="20.335" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.991 22.885v5.63"/><circle cx="37.65" cy="20.335" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.65 22.885v5.63m-11.472-3.506a2.13 2.13 0 0 1 2.124-2.125a2.13 2.13 0 0 1 2.125 2.125v3.4m-4.249-5.524v5.523m4.249-3.399a2.13 2.13 0 0 1 2.125-2.125a2.13 2.13 0 0 1 2.124 2.125v3.4"/>`),
		g.Group(children),
	)
}