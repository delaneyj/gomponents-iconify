package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeStructureDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M72 112v32a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8v-32a8 8 0 0 1 8-8h32a8 8 0 0 1 8 8Zm144-72h-48a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8V48a8 8 0 0 0-8-8Zm0 112h-48a8 8 0 0 0-8 8v48a8 8 0 0 0 8 8h48a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8Z" opacity=".2"/><path d="M168 112h48a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16h-48a16 16 0 0 0-16 16v16h-8a32 32 0 0 0-32 32v24H80v-8a16 16 0 0 0-16-16H32a16 16 0 0 0-16 16v32a16 16 0 0 0 16 16h32a16 16 0 0 0 16-16v-8h32v24a32 32 0 0 0 32 32h8v16a16 16 0 0 0 16 16h48a16 16 0 0 0 16-16v-48a16 16 0 0 0-16-16h-48a16 16 0 0 0-16 16v16h-8a16 16 0 0 1-16-16V96a16 16 0 0 1 16-16h8v16a16 16 0 0 0 16 16ZM64 144H32v-32h32v32Zm104 16h48v48h-48Zm0-112h48v48h-48Z"/></g>`),
		g.Group(children),
	)
}