package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsMNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 36V17.524l6.464 7.756a2 2 0 0 0 3.072 0L32 17.524V36a2 2 0 1 0 4 0V12a2 2 0 0 0-3.536-1.28L24 20.876L15.536 10.72A2 2 0 0 0 12 12v24a2 2 0 1 0 4 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}