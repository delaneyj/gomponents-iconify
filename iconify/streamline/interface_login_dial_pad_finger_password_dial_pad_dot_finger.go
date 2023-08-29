package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLoginDialPadFingerPasswordDialPadDotFinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5V8.25A1.25 1.25 0 0 1 8.25 7h0A1.25 1.25 0 0 1 9.5 8.25V11h2a2 2 0 0 1 2 2v.5"/><circle cx="1" cy="1" r=".5"/><circle cx="5" cy="1" r=".5"/><circle cx="9" cy="1" r=".5"/><circle cx="1" cy="4.5" r=".5"/><circle cx="5" cy="4.5" r=".5"/><circle cx="9" cy="4.5" r=".5"/><circle cx="1" cy="8" r=".5"/><path d="M5 8.5a.5.5 0 0 1 0-1Z"/></g>`),
		g.Group(children),
	)
}