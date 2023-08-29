package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hsl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.773 9.689l10.186 4.031l4.352 10.053l-4.031 10.186l-10.053 4.352l-10.186-4.031l-4.352-10.053l4.031-10.186l10.053-4.352z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.773 9.69c-4.296-1.832-3.188-7.19 0-7.19s4.296 5.358 0 7.19m.454 28.62c4.296 1.832 3.188 7.19 0 7.19s-4.296-5.358 0-7.19M38.31 23.773c1.832-4.296 7.19-3.188 7.19 0s-5.358 4.296-7.19 0m-28.62.454c-1.832 4.296-7.19 3.188-7.19 0s5.358-4.296 7.19 0m24.59 9.732c4.333-1.743 7.338 2.83 5.084 5.083s-6.827-.75-5.084-5.083M13.72 14.041c-4.333 1.743-7.338-2.83-5.084-5.083s6.827.75 5.084 5.083m20.239-.321c-1.743-4.333 2.83-7.338 5.083-5.084s-.75 6.827-5.083 5.084M14.041 34.28c1.743 4.333-2.83 7.338-5.083 5.084s.75-6.827 5.083-5.084"/>`),
		g.Group(children),
	)
}