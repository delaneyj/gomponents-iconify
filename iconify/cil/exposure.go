package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exposure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M456 40H56a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16ZM72 72h345.373L72 417.373Zm368 368H94.627L440 94.627Z"/><path fill="currentColor" d="M336 368v40h32v-40h40v-32h-40v-40h-32v40h-40v32h40zM112 136h112v32H112z"/>`),
		g.Group(children),
	)
}