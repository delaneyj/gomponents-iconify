package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path fill-rule="evenodd" d="M8 5.75a.75.75 0 0 1 .75.75v1.25H10a.75.75 0 0 1 0 1.5H8.75v1.25a.75.75 0 0 1-1.5 0V9.25H6a.75.75 0 0 1 0-1.5h1.25V6.5A.75.75 0 0 1 8 5.75Zm5.25 2.75a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Zm0 6a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Zm-7.28-.53a.75.75 0 0 1 1.06 0l.97.97l.97-.97a.75.75 0 1 1 1.06 1.06l-.97.97l.97.97a.75.75 0 1 1-1.06 1.06L8 17.06l-.97.97a.75.75 0 0 1-1.06-1.06l.97-.97l-.97-.97a.75.75 0 0 1 0-1.06Zm7.28 3.53a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}