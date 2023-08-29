package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YahooBuzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 462q-64 0-109-45T6 309V2h88v307q0 27 19.5 46.5T160 375t46.5-19.5T226 309q0-28-19.5-47T160 243h-44v-88h44q64 0 109 45t45 109q0 63-45 108t-109 45z"/>`),
		g.Group(children),
	)
}