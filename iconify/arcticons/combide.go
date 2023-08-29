package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Combide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.448 27.294h0a2.09 2.09 0 0 1-2.084-2.084v-1.355a2.09 2.09 0 0 1 2.084-2.084h0a2.09 2.09 0 0 1 2.084 2.084v1.355a2.09 2.09 0 0 1-2.084 2.084Zm4.326-3.334a2.09 2.09 0 0 1 2.084-2.085h0a2.09 2.09 0 0 1 2.084 2.084v3.335m-4.168-5.419v5.419m4.168-3.334a2.09 2.09 0 0 1 2.084-2.085h0a2.09 2.09 0 0 1 2.084 2.084v3.335"/><circle cx="37.666" cy="19.27" r=".834" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.666 21.771v5.523m-15.284 3.917a6.945 6.945 0 0 1-5.909 3.249h0A6.944 6.944 0 0 1 9.5 27.545v-7.032a6.944 6.944 0 0 1 6.915-6.973h.059a6.553 6.553 0 0 1 5.964 3.692m9.016 6.623a2.09 2.09 0 0 1 2.084-2.084h0a2.09 2.09 0 0 1 2.084 2.084v1.355a2.09 2.09 0 0 1-2.084 2.084h0a2.09 2.09 0 0 1-2.084-2.084m0 2.084v-8.336"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}