package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHashCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 10.25a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Zm0 5.5a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Z"/><path d="M10.833 6.004a1 1 0 0 1 .914 1.08l-1 12a1 1 0 1 1-1.994-.167l1-12a1 1 0 0 1 1.08-.913Zm5.5 0a1 1 0 0 1 .914 1.08l-1 12a1 1 0 0 1-1.993-.167l1-12a1 1 0 0 1 1.08-.913Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHashCircleFilled0)"/></g>`),
		g.Group(children),
	)
}