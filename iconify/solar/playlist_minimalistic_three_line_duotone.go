package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistMinimalisticThreeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M20 4H3m5 10H3m9-5H3" opacity=".5"/><path d="M16 17.429C16 18.849 14.88 20 13.5 20S11 18.849 11 17.429c0-1.42 1.12-2.572 2.5-2.572S16 16.008 16 17.43Zm0 0V11"/><path stroke-linecap="round" d="m18.675 8.116l-1.755.842c-.341.164-.512.246-.636.373a1.001 1.001 0 0 0-.223.356c-.061.166-.061.355-.061.734c0 .877 0 1.315.191 1.578a1 1 0 0 0 .63.396c.319.058.714-.132 1.504-.51l1.755-.843c.341-.164.512-.246.636-.373a1 1 0 0 0 .223-.356C21 10.149 21 9.958 21 9.58c0-.876 0-1.315-.191-1.578a1 1 0 0 0-.63-.395c-.319-.059-.714.13-1.504.51Z"/></g>`),
		g.Group(children),
	)
}