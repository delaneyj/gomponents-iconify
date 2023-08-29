package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.685 6.659c-.926.309-.906 1.626.03 1.906l7.493 2.242l2.447 7.71c.293.922 1.596.932 1.902.013L18.86 2.62a1 1 0 0 0-1.265-1.265L1.685 6.66Zm3.633.897l11.012-3.67l-3.698 11.096l-1.677-5.284a1 1 0 0 0-.667-.655l-4.97-1.487Z" clip-rule="evenodd"/><path d="m17.767 1.44l1.044 1.077l-8.828 8.543l-1.044-1.077l8.828-8.543Z"/></g>`),
		g.Group(children),
	)
}