package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenVerticalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 3H7.896l2.6 1.5a3 3 0 0 1 1.5 2.598V15h.504a1.5 1.5 0 0 0 1.5-1.5v-4a.5.5 0 0 1 .146-.354l1.708-1.707A.5.5 0 0 0 16 7.086V4.5A1.5 1.5 0 0 0 14.5 3ZM4.016 4.283c-.01.071-.016.145-.016.22v9.398a2 2 0 0 0 1 1.732l3.745 2.162a1.5 1.5 0 0 0 2.25-1.3V7.099a2 2 0 0 0-1-1.732L6.25 3.204a1.501 1.501 0 0 0-2.234 1.08ZM5.539 2H14.5A2.5 2.5 0 0 1 17 4.5v2.586a1.5 1.5 0 0 1-.44 1.06L15 9.706V13.5a2.5 2.5 0 0 1-2.5 2.5h-.505v.496c0 1.924-2.083 3.127-3.75 2.165L4.5 16.499A3 3 0 0 1 3 13.9V4.5a2.457 2.457 0 0 1 .599-1.624A2.494 2.494 0 0 1 5.539 2Z"/>`),
		g.Group(children),
	)
}