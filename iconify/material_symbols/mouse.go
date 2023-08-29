package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.925 0-4.963-2.038T5 15v-4h14v4q0 2.925-2.038 4.963T12 22ZM5 9q0-2.625 1.7-4.588T11 2.075V9H5Zm8 0V2.075q2.6.375 4.3 2.337T19 9h-6Z"/>`),
		g.Group(children),
	)
}