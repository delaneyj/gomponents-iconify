package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M19.582 55.606c.484.178 1.03.297 1.575.297c.849 0 1.697-.297 2.425-.772l30-15.98l.303-.296c.788-.772 1.212-1.723 1.212-2.792c0-1.07-.425-2.08-1.212-2.792l-.303-.297l-30-16.098c-1.091-.832-2.667-1.01-4-.475c-1.516.594-2.485 2.079-2.485 3.683v31.84c0 1.603.97 3.088 2.485 3.682z"/>`),
		g.Group(children),
	)
}