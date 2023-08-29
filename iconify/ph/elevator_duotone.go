package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M184 104v112H72V104Z" opacity=".2"/><path d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-32 80v96h-40v-96Zm-56 96H80v-96h40Zm88 0h-16V104a8 8 0 0 0-8-8H72a8 8 0 0 0-8 8v104H48V48h160v160ZM152 72a8 8 0 0 1-8 8h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}