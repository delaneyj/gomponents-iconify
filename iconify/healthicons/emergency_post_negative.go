package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyPostNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsEmergencyPostNegative0)"><path d="M19.758 31.97a1 1 0 0 1-.728-1.212L19.72 28H17a1 1 0 1 1 0-2h3.22l.81-3.242a1 1 0 1 1 1.94.485L22.28 26H25a1 1 0 1 1 0 2h-3.22l-.81 3.243a1 1 0 0 1-1.212.727Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm43.293 32.952a.995.995 0 0 0 .642-1.303L38 17.962V9h-9v6h7v2H11.938c-1.348 0-2.536.869-2.924 2.138l-4.883 14C3.543 35.063 5.01 37 7.054 37H29.26c1.348 0 2.536-.869 2.923-2.138l4.803-13.733l4.607 10.249L34.5 33.5l-1 2.5l9.794-3.044l-.001-.004Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsEmergencyPostNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}