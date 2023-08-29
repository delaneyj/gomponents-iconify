package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ups(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.96-2.113 16.32-3.605 16.32-16.515V8.145S34.489 4.5 24 4.5S7.68 8.144 7.68 8.144v18.841C7.68 39.895 22.04 41.387 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.32 9.055c-17.854 0-22.974.795-32.64 10.118m3.82 1.27v5.41a3.279 3.279 0 0 0 3.279 3.279h0a3.279 3.279 0 0 0 3.279-3.279v-5.41m0 5.41v3.279m12.842-.733a3.688 3.688 0 0 0 2.697.733h.735A2.17 2.17 0 0 0 36.5 26.96h0a2.17 2.17 0 0 0-2.167-2.172H32.86a2.17 2.17 0 0 1-2.167-2.172h0a2.17 2.17 0 0 1 2.167-2.173h.736a3.688 3.688 0 0 1 2.696.733m-15.009 4.677a3.279 3.279 0 0 0 3.279 3.279h0a3.279 3.279 0 0 0 3.278-3.279v-2.131a3.279 3.279 0 0 0-3.279-3.279h0a3.279 3.279 0 0 0-3.279 3.279m.001-3.279v13.115"/>`),
		g.Group(children),
	)
}