package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingWoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6.009a2.567 2.567 0 1 0 0-5.134a2.567 2.567 0 0 0 0 5.134Zm1.925 14.116l.642-6.417h1.925v-1.925a4.488 4.488 0 0 0-2.023-3.752L12 11.783L9.531 8.031a4.488 4.488 0 0 0-2.023 3.752v1.925h1.925l.642 6.417h3.85Z"/><path d="M9.433 3.442s0 3.208-1.925 3.208m7.059-3.208s0 3.208 1.925 3.208M1 6.125v4m4-2H1m22-2v4m-4-2h4M5 16.039c-2.443.734-4 1.844-4 3.086c0 2.209 4.925 4 11 4s11-1.791 11-4c0-1.242-1.557-2.352-4-3.086"/></g>`),
		g.Group(children),
	)
}