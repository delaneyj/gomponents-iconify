package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoAccessForServiceAnimalThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l23 23M7.052 17s-2.708 0-4.393 2.973C2.125 20.915 2 22.02 2 23.101V24v-4.995a7.464 7.464 0 0 1 6.444-7.394L9 11.534M7.052 17l3.657.671A20 20 0 0 0 14.32 18h.93l.25.25v2.042c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M7.052 17s-1.266 1.224-1.88 3.648c-.137.54-.172 1.099-.172 1.656V24M.5 8v3.5a3 3 0 0 0 3 3h.013M19.5 14V9.5h1a3 3 0 0 0 3-3v-1H19A2.5 2.5 0 0 1 16.5 3a8.597 8.597 0 0 1-2.992 6.523M19.5 14v5.5m0-5.5a6.243 6.243 0 0 1-5.923-4.27l-.07-.207M8.5 0c0 5.5 2.5 8.5 5 9.5l.008.023m0 0A8.583 8.583 0 0 1 11.068 11"/>`),
		g.Group(children),
	)
}