package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeiseOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.656 29.964l1.805 1.503a19.533 19.533 0 0 0 4.31-13.43C43.069 9.215 36.754 2.5 25.328 2.5c-14.634 0-21.15 10.925-21.15 21.85c0 9.423 6.816 21.15 21.15 21.15c12.028 0 16.438-7.417 16.438-7.417l-1.203-1.704s-4.31 6.114-15.235 6.114S9.59 32.67 9.59 24.351c0-9.623 4.912-17.741 15.737-17.741c12.429 0 15.235 8.42 15.335 12.329c0 4.009-.2 7.717-3.007 11.025Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.81 35.677h0a2.384 2.384 0 0 1-1.203-3.207l8.72-19.746a2.384 2.384 0 0 1 3.208-1.203h0a2.384 2.384 0 0 1 1.203 3.208l-8.72 19.846a2.46 2.46 0 0 1-3.208 1.102Zm9.422-.4h0a2.384 2.384 0 0 1-1.203-3.208l4.41-10.024a2.384 2.384 0 0 1 3.208-1.202h0a2.384 2.384 0 0 1 1.203 3.207l-4.41 10.023a2.504 2.504 0 0 1-3.208 1.203Z"/>`),
		g.Group(children),
	)
}