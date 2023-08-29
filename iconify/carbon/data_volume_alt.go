package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataVolumeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 29c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3Zm0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2Z"/><circle cx="20" cy="21" r="2" fill="currentColor"/><path fill="currentColor" d="M15 19c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3Zm0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2Z"/><path fill="currentColor" d="M13 27H4a2.002 2.002 0 0 1-2-2V5a2.002 2.002 0 0 1 2-2h7.586c.53-.002 1.04.21 1.414.586L16.414 7H28a2.002 2.002 0 0 1 2 2v7h-2V9H15.586l-4-4H4v20h9v2Z"/>`),
		g.Group(children),
	)
}