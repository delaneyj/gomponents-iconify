package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AktiaMobileBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.412 29.703l-3.391-10.976L9.5 29.703m1.174-3.705h4.565m3.127-7.074v10.779m0-2.291l5.178-4.85m-3.596 3.368l4.171 3.773m3.097-10.375v8.893a1.512 1.512 0 0 0 1.698 1.482h.51m-3.907-7.855h3.567m8.552 5.744a2.161 2.161 0 0 1-4.322 0V26.22a2.161 2.161 0 0 1 4.322 0m.864 3.483a.919.919 0 0 1-.864-.845v-4.75"/><ellipse cx="31.228" cy="19.371" fill="currentColor" rx=".765" ry=".75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.228 22.591v7.112M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}