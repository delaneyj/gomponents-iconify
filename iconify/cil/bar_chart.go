package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M136 416h88a24.028 24.028 0 0 0 24-24V184a24.028 24.028 0 0 0-24-24h-88a24.028 24.028 0 0 0-24 24v208a24.028 24.028 0 0 0 24 24Zm8-224h72v192h-72ZM424 16h-88a24.028 24.028 0 0 0-24 24v352a24.028 24.028 0 0 0 24 24h88a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24Zm-8 368h-72V48h72Z"/><path fill="currentColor" d="M48 16H16v480h480v-32H48V16z"/>`),
		g.Group(children),
	)
}