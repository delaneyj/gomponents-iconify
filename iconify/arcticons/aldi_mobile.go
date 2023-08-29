package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AldiMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.875 10.499L14.5 27.501m10.297-17.002l-6.375 17.002m10.297-17.002l-6.375 17.002m7.117-12.752h.871M27.867 19h4.039m-5.632 4.251H33.5"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.629 32.977v5.436h2.718m-4.802-5.436v5.436m6.514 0h2.718m-2.718-5.436h2.718m-2.718 2.718h1.767m-1.767-2.718v5.436m-25.282 0v-5.436l2.719 5.436l2.718-5.436v5.436m1.921-1.835c0 1.02.816 1.835 1.767 1.835c1.02 0 1.835-.815 1.835-1.835v-1.834c0-1.02-.815-1.835-1.835-1.835s-1.767.815-1.767 1.835v1.835Zm7.921-.883c.748 0 1.36.612 1.36 1.36s-.612 1.358-1.36 1.358h-2.243v-5.436h2.243c.748 0 1.36.611 1.36 1.359s-.612 1.36-1.36 1.36Zm0 0h-2.243"/>`),
		g.Group(children),
	)
}