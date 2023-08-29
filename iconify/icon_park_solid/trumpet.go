package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trumpet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrumpet0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" d="M32 28h6v11a3 3 0 1 1-6 0V28Z"/><path stroke-linecap="round" d="M29 12h10a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H28"/><path stroke-linejoin="round" d="M14 14h-2a6 6 0 0 0 0 12h2"/><path fill="#fff" d="M14.198 33a.198.198 0 0 1-.198-.198V6.198c0-.11.089-.198.198-.198H17.5C24.956 6 31 12.044 31 19.5S24.956 33 17.5 33h-3.302Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrumpet0)"/>`),
		g.Group(children),
	)
}