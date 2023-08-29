package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilStudioBackdropCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.5 6a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V6Zm12 0h-11v7h11V6Z"/><path d="M5.25 5.5a.5.5 0 0 1 .5-.5h14.5a.5.5 0 0 1 0 1H5.75a.5.5 0 0 1-.5-.5Zm2.24 7.598L6.11 20h13.78l-1.38-6.902l.98-.196l1.38 6.902A1 1 0 0 1 19.89 21H6.11a1 1 0 0 1-.98-1.196l1.38-6.902l.98.196Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilStudioBackdropCircleFilled0)"/></g>`),
		g.Group(children),
	)
}