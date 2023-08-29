package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Splitalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1007.045 257l-88-53l-164 292q78 91 78 209q0 87-43 160.5T673.545 982t-160.5 43t-160.5-43t-116.5-116.5t-43-160.5q0-118 78-209l-164-292l-88 53q-5 3-12.5-5t-5.5-16l66-218q3-9 10.5-14t16.5-2l215 50q8 2 11 12.5t-2 13.5l-100 60l156 279q66-32 138.5-32t139.5 32l156-279l-100-60q-5-3-2-13.5t11-12.5l215-50q9-3 16.5 2t10.5 14l66 218q2 8-5.5 16t-12.5 5zm-494 256q-79 0-135.5 56.5t-56.5 136t56.5 135.5t136 56t135.5-56t56-135.5t-56.5-136t-135.5-56.5z"/>`),
		g.Group(children),
	)
}