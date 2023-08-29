package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.999-.004h2.004V2h-2.004V-.004ZM4.223 2.803L5.64 4.22L4.223 5.637L2.806 4.22l1.417-1.417Zm15.556 0l1.417 1.417l-1.417 1.417l-1.417-1.417l1.417-1.417ZM12 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm-8 6a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm-4.001-1.004h2.004V13H-.001v-2.004Zm22 0h2.004V13h-2.004v-2.004ZM4.223 18.36l1.417 1.417l-1.417 1.418l-1.417-1.418l1.417-1.417Zm15.556 0l1.417 1.417l-1.417 1.417l-1.417-1.417l1.417-1.416ZM11 21.997h2.004V24H11v-2.004Z"/>`),
		g.Group(children),
	)
}