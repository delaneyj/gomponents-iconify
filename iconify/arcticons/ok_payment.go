package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkPayment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.2 7.253a3.136 3.136 0 0 1 6.22.803M37.642 25.05a2.9 2.9 0 0 1-.148 2.83a3.136 3.136 0 0 1-5.431 0M31.2 5.838a3.136 3.136 0 0 1 5.137 3.597m2.403 26.473a11.563 11.563 0 0 1-10.123 7.532a12.202 12.202 0 0 1-11.438-18.788"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".971" d="M29.74 19.32a10.611 10.611 0 0 1 7.9 5.73"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.703 34.184a3.335 3.335 0 0 1 2.334-2.248a3.136 3.136 0 0 1 3.704 3.972m-6.038-1.724c-1.129 2.42-3.627 3.319-6.246 2.798a5.796 5.796 0 0 1-.491-11.25c2.564-.746 4.761-.164 6.096 2.148M31.2 5.838L17.18 24.651m12.56-5.331l6.596-9.885M15.425 29.021L19.2 7.253m5.125 7.81l1.095-7.008M9.397 18.908a3.136 3.136 0 1 1 5.721-2.57m.692 18.672L9.398 18.909m7.492 1.667l-1.772-4.239"/>`),
		g.Group(children),
	)
}