package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraNoflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 6.5h8a2 2 0 0 1 2 2v6c0 .559-.229 1.064-.598 1.427M15.5 16.5h-11a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1m-2-2l14 14"/><path fill="currentColor" d="M17 9a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.215 9.557a3 3 0 0 0 4.27 4.193M13.5 11.5a3 3 0 0 0-3-3m-1-4h2a1 1 0 0 1 1 1v1h-4v-1a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}