package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BootVolumeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 30h-7c-1.103 0-2-.897-2-2v-2h2v2h7v-7h-2v-2h2c1.103 0 2 .897 2 2v7c0 1.103-.897 2-2 2Z"/><path fill="currentColor" d="M12 24h2v-6.586L22.582 26L24 24.586L15.414 16H22v-2H12v10z"/><path fill="currentColor" d="M8 27H4a2.002 2.002 0 0 1-2-2V5a2.002 2.002 0 0 1 2-2h7.586c.53-.002 1.04.21 1.414.586L16.414 7H28a2.002 2.002 0 0 1 2 2v7h-2V9H15.586l-4-4H4v20h4v2Z"/>`),
		g.Group(children),
	)
}