package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPrinter0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M37 32H11v12h26V32Z"/><path stroke-linecap="round" d="M4 20h40v18h-6.983v-6H10.98v6H4V20Z" clip-rule="evenodd"/><path fill="#555" d="M38 4H10v16h28V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPrinter0)"/>`),
		g.Group(children),
	)
}