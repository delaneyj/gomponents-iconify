package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paydirekt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.328 22.153a3.375 3.375 0 0 1-3.375 3.375h0a3.375 3.375 0 0 1-3.375-3.375v-2.194a3.375 3.375 0 0 1 3.375-3.375h0a3.375 3.375 0 0 1 3.375 3.375m0 5.569v-8.944m7.017 8.944l-3.595-8.944m6.75 0l-4.286 12.149a3.227 3.227 0 0 1-1.91 2.225M11.5 22.153a3.375 3.375 0 0 0 3.375 3.375h0a3.375 3.375 0 0 0 3.375-3.375v-2.194a3.375 3.375 0 0 0-3.375-3.375h0A3.375 3.375 0 0 0 11.5 19.96m0-3.376v13.5"/><circle cx="23.656" cy="33.509" r="1.907" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.831" cy="32.259" r="1.301" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.604" cy="32.657" r="1.301" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.625" cy="30.834" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}