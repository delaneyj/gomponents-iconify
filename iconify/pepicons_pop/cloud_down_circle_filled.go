package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCloudDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M14 6h-1a4.002 4.002 0 0 0-3.874 3H9a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 14 6Zm-3.324 5l.387-1.501A2.002 2.002 0 0 1 13 8h1c.937 0 1.743.65 1.95 1.549l.28 1.221l1.221.28A2.002 2.002 0 0 1 17 15H9a2 2 0 1 1 0-4h1.676Z" clip-rule="evenodd"/><path d="M12.5 13a1 1 0 1 1 2 0v7.5a1 1 0 1 1-2 0V13Z"/><path d="M15.375 17.72a1 1 0 1 1 1.25 1.56l-2.5 2a1 1 0 0 1-1.25-1.56l2.5-2Z"/><path d="M10.375 19.28a1 1 0 1 1 1.25-1.56l2.5 2a1 1 0 0 1-1.25 1.56l-2.5-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCloudDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}