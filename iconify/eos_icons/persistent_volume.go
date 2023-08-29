package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersistentVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 1H10a2 2 0 0 0-2 2v10.368a6 6 0 0 1 .989.455A5.983 5.983 0 0 1 17.65 21H20a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2ZM9 3a1 1 0 1 1 1 1a1 1 0 0 1-1-1Zm6 10a4.973 4.973 0 0 1-2-.422V10h-2.578A4.997 4.997 0 1 1 15 13Zm5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0-16a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M12 23a3.973 3.973 0 0 1-3-1.38a3.967 3.967 0 0 1-1-2.617V19a2 2 0 1 0-2 2a1.928 1.928 0 0 0 .79-.175a5.95 5.95 0 0 0 .927 1.765A3.834 3.834 0 0 1 6 23a4 4 0 1 1 3-6.62a3.967 3.967 0 0 1 1 2.617V19a2 2 0 1 0 2-2a1.928 1.928 0 0 0-.79.176a5.95 5.95 0 0 0-.928-1.766A3.834 3.834 0 0 1 12 15a4 4 0 0 1 0 8Z"/><circle cx="15" cy="8" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}