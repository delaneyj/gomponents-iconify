package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetrinaAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsLetrinaAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM21 19a1 1 0 1 0 0 2h20a1 1 0 1 0 0-2H21Zm-9 6H6V12h12v11h24c0 8.43-5.703 13.35-13 13.94V37h-5.5v2H29v3H14.5V30c-.747 0-2.5-.5-2.5-2v-3ZM6 8a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v2H6V8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsLetrinaAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}