package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPrinterTwo0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M38 20V8a2 2 0 0 0-2-2H12a2 2 0 0 0-2 2v12"/><rect width="36" height="22" x="6" y="20" rx="2"/><path fill="#fff" stroke-linecap="round" stroke-linejoin="round" d="M20 34h15v8H20v-8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 26h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPrinterTwo0)"/>`),
		g.Group(children),
	)
}