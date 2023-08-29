package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m6.804 2.632l-.637.264A3.507 3.507 0 0 0 4 6.137v4.386a1.46 1.46 0 0 0 2.167 1.276l.227-.126a4 4 0 0 1 3.88 0l.453.251a4 4 0 0 0 3.88 0l.734-.407A3.222 3.222 0 0 0 17 8.7V4.638a1.605 1.605 0 0 0-2.07-1.534l-.893.272a4 4 0 0 1-2.694-.13l-1.48-.614a4 4 0 0 0-3.059 0ZM5 6.137c0-1.014.611-1.929 1.549-2.317l.638-.264a3 3 0 0 1 2.293 0l1.481.613a5 5 0 0 0 3.367.163l.893-.271a.604.604 0 0 1 .779.577V8.7c0 .807-.438 1.551-1.144 1.943l-.735.407a3 3 0 0 1-2.91 0l-.453-.252a5 5 0 0 0-4.85 0l-.226.126A.46.46 0 0 1 5 10.523V6.137Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 2a1 1 0 0 1 1 1v14a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}