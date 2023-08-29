package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeBinding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19a6 6 0 0 1 6-6h4V3a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16.001A2 2 0 0 0 4 21h2.349A5.976 5.976 0 0 1 6 19Zm8-17a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM9 3a5 5 0 1 1-2 9.578V10H4.422A4.991 4.991 0 0 1 9 3ZM4 2a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm0 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><circle cx="9" cy="8" r="2" fill="currentColor"/><path fill="currentColor" d="M13 18h4v2h-4z"/><path fill="currentColor" d="M18 15h-2v2h2a2 2 0 0 1 0 4h-2v2h2a4 4 0 0 0 0-8Zm-4 6h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Z"/>`),
		g.Group(children),
	)
}