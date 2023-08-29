package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtalkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.03 35.763l3.442 2.385l-4.255 3.496C-8.375 21.47 7.689 10.004 16.34 7.386c3.787-1.145 15.857-2.864 23.44 4.533c8.353 8.148 2.44 21.423-15.75 23.844Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.902 12.605c-5.688-.47-5.908 14.716-11.03 14.174c-5.083-.537-5.11-14.193 0-14.174c5.07.021 3.22 12.811 11.03 14.174m-2.07-10.47h4.487m-2.202 6.772v-6.772m5.375 5.078A1.698 1.698 0 0 1 28.8 23.08h0a1.698 1.698 0 0 1-1.693-1.693v-1.1a1.698 1.698 0 0 1 1.693-1.693h0a1.698 1.698 0 0 1 1.693 1.693m-.001 2.793v-4.486m2.037-2.285v6.772m2.035-6.773v6.772m0-1.439l3.048-3.047m-2.117 2.116l2.456 2.37"/>`),
		g.Group(children),
	)
}