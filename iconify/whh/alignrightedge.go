package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignrightedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.5 1024q-26.5 0-45-19T896 960V64q0-26 18.5-45t45-19t45.5 19t19 45v896q0 27-19 45.5t-45.5 18.5zM704 960H64q-27 0-45.5-18.5T0 896V640q0-26 18.5-45T64 576h640q26 0 45 19t19 45v256q0 27-19 45.5T704 960zm0-512H320q-27 0-45.5-18.5T256 384V128q0-26 18.5-45T320 64h384q26 0 45 19t19 45v256q0 27-19 45.5T704 448zm-64-224q0-13-9.5-22.5T608 192H416q-13 0-22.5 9.5T384 224v64q0 13 9.5 22.5T416 320h192q13 0 22.5-9.5T640 288v-64z"/>`),
		g.Group(children),
	)
}