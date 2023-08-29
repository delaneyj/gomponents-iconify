package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpreadsheetsNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSpreadsheetsNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM30 4H11a1 1 0 0 0-1 1v38a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V12h-7a1 1 0 0 1-1-1V4ZM14 19a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H15a1 1 0 0 1-1-1V19Zm2 1v3h16v-3H16Zm0 8v-3h4v3h-4Zm0 5v-3h4v3h-4Zm6-3v3h4v-3h-4Zm6 3v-3h4v3h-4Zm0-5v-3h4v3h-4Zm-6 0h4v-3h-4v3Zm10-18h5.414L32 4.586V10Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSpreadsheetsNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}