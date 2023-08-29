package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.626 27.902c.63.755 1.384 1.133 2.517 1.133h1.51c1.385 0 2.518-1.133 2.518-2.518S15.038 24 13.653 24h-1.636C10.633 24 9.5 22.867 9.5 21.483s1.133-2.517 2.517-2.517h1.51c1.133 0 1.888.251 2.518 1.133m7.251 7.677c-.377.755-1.258 1.259-2.14 1.259a2.525 2.525 0 0 1-2.517-2.518v-1.636c0-1.384 1.133-2.517 2.518-2.517s2.517 1.133 2.517 2.517v.881h-5.035m7.504-6.797v8.81c0 .756.503 1.26 1.258 1.26h.378m2.316-10.07v8.81c0 .756.503 1.26 1.259 1.26h.377m6.769-2.518c0 1.385-1.133 2.517-2.517 2.517s-2.517-1.132-2.517-2.517v-1.636c0-1.385 1.132-2.517 2.517-2.517S38.5 23.497 38.5 24.88m0 4.155v-6.797"/>`),
		g.Group(children),
	)
}