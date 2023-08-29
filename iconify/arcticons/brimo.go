package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brimo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M34.5 14.377v11.232m-10.51 0V14.377h3.677c2.079 0 3.764 1.69 3.764 3.772s-1.685 3.772-3.764 3.772H23.99m3.677 0l3.677 3.685"/><rect width="4.468" height="5.921" x="30.032" y="27.702" rx="2.234" ry="2.234"/><path d="M18.693 33.623v-3.687c0-1.234 1-2.234 2.234-2.234h0c1.234 0 2.235 1 2.235 2.234v3.687m0-3.687c0-1.234 1-2.234 2.234-2.234h0c1.234 0 2.234 1 2.234 2.234v3.687m-9.497-13.63a2.808 2.808 0 0 1 0 5.616H13.5V14.378h4.633a2.808 2.808 0 0 1 0 5.615h0Zm0 0H13.5"/></g>`),
		g.Group(children),
	)
}