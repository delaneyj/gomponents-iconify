package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kijiji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.482 20.09v8.581a2.44 2.44 0 0 1-2.429 2.452h0"/><circle cx="17.314" cy="17.933" r=".75" fill="currentColor"/><circle cx="22.482" cy="17.933" r=".75" fill="currentColor"/><circle cx="27.65" cy="17.933" r=".75" fill="currentColor"/><circle cx="32.818" cy="17.933" r=".75" fill="currentColor"/><circle cx="37.986" cy="17.933" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.314 20.09v6.497M27.65 20.09v6.497m5.168-6.497v8.172a2.44 2.44 0 0 1-2.429 2.452h0m7.597-10.624v6.497m-28.722-9.71v9.807m.001-2.084l4.439-4.417m-3.026 3.012l3.49 3.474"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}