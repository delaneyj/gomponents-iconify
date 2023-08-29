package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batterybot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.505v6.67a1 1 0 0 1-1 1h-3v4.63a2 2 0 0 1-2 2h-31a2 2 0 0 1-2-2v-17.61a2 2 0 0 1 2-2h31a2 2 0 0 1 2 2v4.31h3a1 1 0 0 1 1 1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".986" d="M18 13.275v21.027"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 19.505v8.67"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".973" d="M11.293 13.574v20.445"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".989" d="M32.344 13.285V34.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".987" d="M25.225 13.559v21.043"/>`),
		g.Group(children),
	)
}