package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFlagCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="m9.804 5.632l-.637.264A3.507 3.507 0 0 0 7 9.137v4.386a1.46 1.46 0 0 0 2.167 1.276l.227-.126a4 4 0 0 1 3.88 0l.453.251a4 4 0 0 0 3.88 0l.733-.407A3.222 3.222 0 0 0 20 11.7V7.638a1.605 1.605 0 0 0-2.07-1.534l-.893.272a4 4 0 0 1-2.693-.13l-1.482-.614a4 4 0 0 0-3.058 0ZM8 9.137c0-1.014.611-1.929 1.549-2.317l.638-.264a3 3 0 0 1 2.293 0l1.481.613a5 5 0 0 0 3.367.163l.893-.271a.604.604 0 0 1 .779.577V11.7c0 .807-.438 1.551-1.144 1.943l-.735.407a3 3 0 0 1-2.91 0l-.453-.252a5 5 0 0 0-4.85 0l-.226.126A.46.46 0 0 1 8 13.523V9.137Z"/><path d="M8 5a1 1 0 0 1 1 1v14a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFlagCircleFilled0)"/></g>`),
		g.Group(children),
	)
}