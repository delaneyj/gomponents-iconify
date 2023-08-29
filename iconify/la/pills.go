package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pills(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.5 5A5.506 5.506 0 0 0 3 10.5v11C3 24.532 5.467 27 8.5 27s5.5-2.468 5.5-5.5v-11C14 7.468 11.533 5 8.5 5zm0 2c1.93 0 3.5 1.57 3.5 3.5V15H5v-4.5C5 8.57 6.57 7 8.5 7zm14 7a6.508 6.508 0 0 0-6.5 6.5c0 3.584 2.916 6.5 6.5 6.5s6.5-2.916 6.5-6.5s-2.916-6.5-6.5-6.5zm0 2c2.481 0 4.5 2.019 4.5 4.5c0 .879-.262 1.693-.7 2.387l-6.187-6.188A4.459 4.459 0 0 1 22.5 16zM5 17h7v4.5c0 1.93-1.57 3.5-3.5 3.5S5 23.43 5 21.5V17zm13.7 1.113l6.187 6.188A4.462 4.462 0 0 1 22.5 25a4.505 4.505 0 0 1-4.5-4.5c0-.879.262-1.693.7-2.387z"/>`),
		g.Group(children),
	)
}