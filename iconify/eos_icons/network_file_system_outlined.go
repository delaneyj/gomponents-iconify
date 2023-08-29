package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkFileSystemOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17h-3v-4h-2v4H8v1H2v2h6v1h8v-1h6v-2h-6v-1z"/><path fill="currentColor" d="M18 4h-5.793l-.853-.854L11.207 3H6a1.502 1.502 0 0 0-1.5 1.5v8A1.502 1.502 0 0 0 6 14h12a1.502 1.502 0 0 0 1.5-1.5v-7A1.502 1.502 0 0 0 18 4Zm-.5 2v6h-11V5h3.793l.853.854l.147.146H17.5Z"/>`),
		g.Group(children),
	)
}