package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyringeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.678 7.98l-1.414 1.413l-2.122-2.121l-2.121 2.121l3.536 3.536l-1.415 1.414l-.707-.707L11.071 20H5.414l-2.12 2.121l-1.415-1.414L4 18.586v-5.657l6.364-6.364l-.707-.707l1.414-1.414l3.536 3.535l2.121-2.121l-2.121-2.122l1.414-1.414l5.657 5.657Zm-12.02 6.363l-2.83-2.828l-1.414 1.414l2.829 2.828l1.414-1.414Zm2.828-2.828l-2.83-2.829L8.243 10.1l2.828 2.829l1.415-1.414Z"/>`),
		g.Group(children),
	)
}