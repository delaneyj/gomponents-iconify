package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedwoodCu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.832 5.2a15.739 15.739 0 0 1 5.913-1.145h8.51c1.952 0 3.82.354 5.544 1.001l.09 8.542l8.8-.125a15.61 15.61 0 0 1 1.31 6.284v6.982s-1.689-4.503-16.534-9.673C15.775 12.994 13.012 5.532 13.832 5.2Zm27.77 26.014s-1.192-5.407-15.962-10.46c-14.885-5.089-14.521-12.33-14.521-12.33M39.702 36.95s-1.51-8.165-15.99-12.815C10.82 19.994 9.202 15.229 8.398 12.354m25.649 30.569c-1.88.79-3.624 1.022-5.792 1.022h-8.51a15.74 15.74 0 0 1-6.225-1.275l.012-8.288l-8.286-.007A15.615 15.615 0 0 1 4 28.244v-7.115s1.68 4.238 14.46 8.636C33.18 34.83 34.67 42.66 34.048 42.923Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.944 40.453s-1.279-8.539-14.87-12.984C7.54 22.714 6.3 16.994 6.3 16.994"/>`),
		g.Group(children),
	)
}