package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassFullFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m5.004 10.229l-.003-.186l.001-.113l.008-.071l1-7a1 1 0 0 1 .877-.853L7 2h10a1 1 0 0 1 .968.747l.022.112l1.006 7.05L19 10c0 3.226-2.56 5.564-6 5.945V20h3a1 1 0 0 1 .117 1.993L16 22H8a1 1 0 0 1-.117-1.993L8 20h3v-4.055c-3.358-.371-5.878-2.609-5.996-5.716zM16.133 4H7.866l-.607 4.258a6.001 6.001 0 0 1 5.125.787l.216.155a4 4 0 0 0 4.32.31L16.133 4z"/></g>`),
		g.Group(children),
	)
}