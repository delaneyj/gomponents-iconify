package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 80H40a24.028 24.028 0 0 0-24 24v240a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24Zm-8 256H48V112h416Z"/><path fill="currentColor" d="M144 272h160v32H144zm-64 0h32v32H80zm320 0h32v32h-32zm-64 0h32v32h-32zm32-64h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm224-64h32v32h-32zm64 0h32v32h-32zm-128 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32H80z"/>`),
		g.Group(children),
	)
}