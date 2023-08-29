package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v384a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24Zm-8 400H48V72h416Z"/><path fill="currentColor" d="M152 240h32v-40h40v-32h-40v-40h-32v40h-40v32h40v40zm44.284 45.089L168 313.373l-28.284-28.284l-22.627 22.627L145.373 336l-28.284 28.284l22.627 22.627L168 358.627l28.284 28.284l22.627-22.627L190.627 336l28.284-28.284l-22.627-22.627zM288 168h112v32H288zm0 120h112v32H288zm0 64h112v32H288z"/>`),
		g.Group(children),
	)
}