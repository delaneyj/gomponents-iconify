package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.685 9.659c-.926.309-.906 1.626.03 1.906l7.493 2.242l2.447 7.71c.293.922 1.596.932 1.902.013L21.86 5.62a1 1 0 0 0-1.265-1.265L4.685 9.66Zm3.633.897l11.012-3.67l-3.698 11.096l-1.677-5.284a1 1 0 0 0-.667-.655l-4.97-1.487Z" clip-rule="evenodd"/><path d="m20.767 4.44l1.044 1.077l-8.828 8.543l-1.044-1.077l8.828-8.543Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}