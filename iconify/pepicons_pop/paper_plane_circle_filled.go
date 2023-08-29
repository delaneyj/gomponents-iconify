package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPaperPlaneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M4.685 9.659c-.926.309-.906 1.626.03 1.906l7.493 2.242l2.447 7.71c.293.922 1.596.932 1.902.013L21.86 5.62a1 1 0 0 0-1.265-1.265L4.685 9.66Zm3.633.897l11.012-3.67l-3.698 11.096l-1.677-5.284a1 1 0 0 0-.667-.655l-4.97-1.487Z" clip-rule="evenodd"/><path d="m20.767 4.44l1.044 1.077l-8.828 8.543l-1.044-1.077l8.828-8.543Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPaperPlaneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}