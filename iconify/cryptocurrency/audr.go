package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-5.146-18.736l3.61 1.92l1.92-.96l-4.589-2.486V7.36l8.823 4.58l1.92-.96l-11.3-5.924A.845.845 0 0 0 10 5.805v6.067c.003.587.332 1.124.854 1.392zm13.757-1.018a.845.845 0 0 0-.96-.153l-12.797 6.643c-.525.27-.855.811-.854 1.402v6.057a.845.845 0 0 0 1.238.749l12.768-6.624c.533-.276.864-.83.855-1.43v-6.048a.835.835 0 0 0-.25-.596zm-1.536 6.538l-11.28 5.856v-4.378l11.28-5.856v4.378z"/>`),
		g.Group(children),
	)
}