package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscountSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.448.436l-1.13 1.129a.488.488 0 0 1-.344.143H3.196c-.822 0-1.488.666-1.488 1.488v1.778a.49.49 0 0 1-.143.345L.435 6.448a1.488 1.488 0 0 0 0 2.104l1.13 1.13a.488.488 0 0 1 .143.344v1.778c0 .822.666 1.488 1.488 1.488h1.778a.49.49 0 0 1 .345.143l1.129 1.13a1.488 1.488 0 0 0 2.104 0l1.13-1.13a.488.488 0 0 1 .344-.143h1.778c.822 0 1.488-.666 1.488-1.488v-1.778a.49.49 0 0 1 .143-.345l1.13-1.129a1.488 1.488 0 0 0 0-2.104l-1.13-1.13a.488.488 0 0 1-.143-.344V3.196c0-.822-.666-1.488-1.488-1.488h-1.778a.488.488 0 0 1-.345-.143L8.552.435a1.488 1.488 0 0 0-2.104 0Zm-1.802 9.21l5-5l.708.708l-5 5l-.708-.708ZM5 5v1h1V5H5Zm4 5h1V9H9v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}