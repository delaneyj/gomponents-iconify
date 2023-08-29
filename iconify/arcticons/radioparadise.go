package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioparadise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="23.188" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.943" ry="15.687" transform="rotate(-19.915 24 23.188)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.092 26.95a6.208 6.208 0 0 0 7.517 4.487h0a6.208 6.208 0 0 0 4.487-7.517l-.985-3.901a6.208 6.208 0 0 0-7.517-4.487h0a6.208 6.208 0 0 0-4.487 7.517m-1.515-6.003l6.06 24.008m-14.028-20.65a6.371 6.371 0 0 1 4.578-7.73h0m-6.154 1.575l4.176 16.309"/>`),
		g.Group(children),
	)
}