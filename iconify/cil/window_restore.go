package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M352 153H40.247a24.028 24.028 0 0 0-24 24v281a24.028 24.028 0 0 0 24 24H352a24.028 24.028 0 0 0 24-24V177a24.028 24.028 0 0 0-24-24Zm-8 32v45.22H48.247V185ZM48.247 450V262.22H344V450Z"/><path fill="currentColor" d="M472 32H152a24.028 24.028 0 0 0-24 24v65h32V64h304v275.143h-56v32h64a24.028 24.028 0 0 0 24-24V56a24.028 24.028 0 0 0-24-24Z"/>`),
		g.Group(children),
	)
}