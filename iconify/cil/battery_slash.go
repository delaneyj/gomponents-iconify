package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatterySlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H163.882l32 32H400v80h64v104h-48v32h80V176Zm-281.373-48l-32-32l-80-80H16v22.627L73.373 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59h361.387l72 72H496v-22.627L266.563 243.937ZM47.986 392V128h57.387l264 264Z"/>`),
		g.Group(children),
	)
}