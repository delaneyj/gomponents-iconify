package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlaceholderRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 6h12V4a2 2 0 0 0-2-2h-6z"/><rect width="20" height="14" y="4" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}