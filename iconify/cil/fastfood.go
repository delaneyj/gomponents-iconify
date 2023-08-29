package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastfood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M474.183 120H360V40h-32v80H200v32h17.92l6.154 48H112a80.091 80.091 0 0 0-80 80v16h280v-16a79.508 79.508 0 0 0-8-34.846a80.248 80.248 0 0 0-47.155-41.185L250.183 152h187.634l-35.9 280H32v32h398.08l40-312H496v-32ZM277.258 264H66.742A48.083 48.083 0 0 1 112 232h120a48.083 48.083 0 0 1 45.258 32Z"/><path fill="currentColor" d="M304 352h8v-32H32v32h272zm0 56h8v-32H32v32h272z"/>`),
		g.Group(children),
	)
}