package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Industry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M459.26 96L328 225.7V96h-34.525L168 223.267V16H16v480h480V96ZM464 464H48V48h88v216h36.778L296 139.018V264h38.764L464 136.3Z"/><path fill="currentColor" d="M136 328v8h32v-32h-32v24zm0 48h32v32h-32zm80-48v8h32v-32h-32v24zm0 48h32v32h-32zm80-48v8h32v-32h-32v24zm0 48h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32z"/>`),
		g.Group(children),
	)
}