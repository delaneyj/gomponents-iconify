package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudioDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m72 160l24-24v88l-24-24H48v-40Zm80-128v56h56Z" opacity=".2"/><path d="M99.06 128.61a8 8 0 0 0-8.72 1.73L68.69 152H48a8 8 0 0 0-8 8v40a8 8 0 0 0 8 8h20.69l21.65 21.66A8 8 0 0 0 104 224v-88a8 8 0 0 0-4.94-7.39ZM88 204.69l-10.34-10.35A8 8 0 0 0 72 192H56v-24h16a8 8 0 0 0 5.66-2.34L88 155.31ZM152 180a40.55 40.55 0 0 1-20 34.91a8 8 0 0 1-8-13.82a24.49 24.49 0 0 0 0-42.18a8 8 0 0 1 8-13.82A40.55 40.55 0 0 1 152 180Zm61.66-97.66l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v80a8 8 0 0 0 16 0V40h88v48a8 8 0 0 0 8 8h48v120h-32a8 8 0 0 0 0 16h32a16 16 0 0 0 16-16V88a8 8 0 0 0-2.34-5.66ZM160 51.31L188.69 80H160Z"/></g>`),
		g.Group(children),
	)
}