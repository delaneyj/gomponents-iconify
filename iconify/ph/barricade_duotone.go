package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 80v68l-76-76h68a8 8 0 0 1 8 8ZM32 72a8 8 0 0 0-8 8v4l76 76h72L84 72Z" opacity=".2"/><path d="M224 64H32a16 16 0 0 0-16 16v72a16 16 0 0 0 16 16h24v32a8 8 0 0 0 16 0v-32h112v32a8 8 0 0 0 16 0v-32h24a16 16 0 0 0 16-16V80a16 16 0 0 0-16-16Zm0 64.69L175.31 80H224ZM80.69 80l72 72h-49.38L32 80.69V80ZM32 103.31L80.69 152H32ZM224 152h-48.69l-72-72h49.38L224 151.32v.68Z"/></g>`),
		g.Group(children),
	)
}