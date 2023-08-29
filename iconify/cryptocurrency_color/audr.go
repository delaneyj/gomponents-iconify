package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#34318A"/><path fill="#FFF" d="m10.854 13.264l3.61 1.92l1.92-.96l-4.589-2.486V7.36l8.823 4.58l1.92-.96l-11.3-5.924A.845.845 0 0 0 10 5.805v6.067c.003.587.332 1.124.854 1.392zm13.757-1.018a.845.845 0 0 0-.96-.153l-12.797 6.643c-.525.27-.855.811-.854 1.402v6.057a.845.845 0 0 0 1.238.749l12.768-6.624c.533-.276.864-.83.855-1.43v-6.048a.835.835 0 0 0-.25-.596zm-1.536 6.538l-11.28 5.856v-4.378l11.28-5.856v4.378z"/></g>`),
		g.Group(children),
	)
}