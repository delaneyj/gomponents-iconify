package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M39.986 424H408a23.825 23.825 0 0 0 24-23.59V344h64V176h-64v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59Zm8-296H400v80h64v104h-64v80H47.986Z"/>`),
		g.Group(children),
	)
}