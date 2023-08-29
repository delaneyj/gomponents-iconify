package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.41 16.192l1.18 1.615L10 14.584V21a1 1 0 0 0 1.541.841l7-4.5a.999.999 0 0 0 .049-1.649L13.537 12l5.053-3.692a1.002 1.002 0 0 0-.049-1.65l-7-4.5a1.002 1.002 0 0 0-1.021-.037c-.32.176-.52.513-.52.879v6.416L5.59 6.192L4.41 7.808L10 11.893v.215l-5.59 4.084zM12 4.832l4.232 2.721L12 10.646V4.832zm0 8.522l4.232 3.093L12 19.168v-5.814z"/>`),
		g.Group(children),
	)
}