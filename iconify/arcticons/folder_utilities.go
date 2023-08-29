package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderUtilities(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.678 13.044H24.77c-1.963-.108-5.931-4.238-8.188-4.238H6.68v-.001a2.176 2.176 0 0 0-2.18 2.171v7.307h39v-3.418a1.822 1.822 0 0 0-1.822-1.821Zm1.822 5.239h-39v18.733a2.176 2.176 0 0 0 2.174 2.18h34.645a2.176 2.176 0 0 0 2.181-2.172V18.283Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.278 25.858l1.8 1.471l2.175-.823l.374-2.295l-1.8-1.472m6.891 8.233l-5.465-4.466m9.44 5.114l-1.801-1.471l-2.174.823l-.374 2.295l1.8 1.472"/>`),
		g.Group(children),
	)
}