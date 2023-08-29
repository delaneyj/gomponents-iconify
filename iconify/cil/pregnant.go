package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pregnant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 144a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64Zm0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32Zm129.959 203.37c-15.021-16.9-35.063-27.659-62.61-33.506L266.551 160h-88.428L152 342.863V400h56v96h96v-96h80v-48c0-44.972-9.826-77.888-30.041-100.63ZM352 368h-80v96h-32v-96h-56v-22.863L205.877 192h39.572l23.291 54.344l8.629 1.438c24.5 4.083 41.233 11.979 52.672 24.848C344.817 289.253 352 315.215 352 352Z"/>`),
		g.Group(children),
	)
}