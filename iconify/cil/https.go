package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Https(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400 200h-12v-56a128 128 0 0 0-256 0v56h-12a24.028 24.028 0 0 0-24 24v248a24.028 24.028 0 0 0 24 24h280a24.028 24.028 0 0 0 24-24V224a24.028 24.028 0 0 0-24-24Zm-236-56a96 96 0 0 1 192 0v56H164Zm228 320H128V232h264Z"/><path fill="currentColor" d="M240 328h40v40h-40z"/>`),
		g.Group(children),
	)
}