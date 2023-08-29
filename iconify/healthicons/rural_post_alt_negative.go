package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuralPostAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsRuralPostAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6.172 21.102A4.858 4.858 0 0 0 8 26.081V42h4V29h8v13h20V26.077a4.858 4.858 0 0 0 1.824-4.976l-1.114-4.875a1.517 1.517 0 0 0-1.479-1.179H31v-5a4 4 0 0 0-4-4h-6a4 4 0 0 0-4 4v5H8.764c-.707 0-1.321.49-1.479 1.18l-1.113 4.874Zm10.954-4.054h-1.177v4.52a3.481 3.481 0 0 0 6.962 0v-1.52H21a4.002 4.002 0 0 1-3.874-3Zm7.958 3v1.52a3.481 3.481 0 0 0 6.962 0v-4.52h-1.172a4.002 4.002 0 0 1-3.874 3h-1.916Zm.916 11h9v6h-9v-6Zm-2-2h13v10H24v-10Zm-14.85-12l-1.028 4.5a2.863 2.863 0 1 0 5.654.637v-5.137H9.15Zm29.696 0H34.22v5.137a2.863 2.863 0 1 0 5.655-.637l-1.028-4.5ZM24 8.548a1 1 0 0 1 1 1v2.5h2.5a1 1 0 1 1 0 2H25v2.5a1 1 0 1 1-2 0v-2.5h-2.5a1 1 0 1 1 0-2H23v-2.5a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRuralPostAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}