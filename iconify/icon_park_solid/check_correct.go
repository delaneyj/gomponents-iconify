package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCorrect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCheckCorrect0"><g fill="none"><g stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-path="url(#ipSCheckCorrect1)"><path d="M42 20v19a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h21"/><path d="m16 20l10 8L41 7"/></g><defs><clipPath id="ipSCheckCorrect1"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCheckCorrect0)"/>`),
		g.Group(children),
	)
}