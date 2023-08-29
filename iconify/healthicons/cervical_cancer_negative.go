package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CervicalCancerNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsCervicalCancerNegative0)"><path fill-rule="evenodd" d="M26.735 14.455L25.152 16.3l3.77 4.398l-1.274 1.204L24 17.645l-3.649 4.257l-1.274-1.204l3.77-4.398l-1.581-1.845c-1.052-1.226-1.005-3.102.107-4.269A1.861 1.861 0 0 1 22.72 9.6h2.56c.505 0 .99.21 1.348.586c1.11 1.167 1.158 3.043.107 4.27Zm-4.15-2.996a.187.187 0 0 1 .135-.059h2.56c.05 0 .1.021.135.059a1.35 1.35 0 0 1 .045 1.792L24 14.955l-1.46-1.704a1.35 1.35 0 0 1 .045-1.792Z" clip-rule="evenodd"/><path d="m31.88 24.173l-7.023 7.373V42h-1.714V31.546l-7.022-7.373l1.212-1.273l6.667 7l6.667-7l1.212 1.273Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM10.705 6h26.59c3.126 7.848 4.713 14.106 4.705 19.852c-.008 5.614-1.538 10.66-4.504 16.148H10.504C7.538 36.511 6.008 31.466 6 25.852C5.992 20.105 7.579 13.848 10.705 6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCervicalCancerNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}