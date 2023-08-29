package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeBindingOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18h4v2h-4z"/><path fill="currentColor" d="M19 15h-2v2h2a2 2 0 0 1 0 4h-2v2h2a4 4 0 0 0 0-8Zm-6 2h2v-2h-2a4 4 0 0 0 0 8h2v-2h-2a2 2 0 0 1 0-4ZM6 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1ZM6 20a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/><path fill="currentColor" d="M9.69 24H3a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v11h-2V2H3v20h4.812a6.04 6.04 0 0 0 1.879 2Z"/><path fill="currentColor" d="M14.578 13A4.998 4.998 0 1 0 6 14h3v.54A5.969 5.969 0 0 1 13 13ZM10 13a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}