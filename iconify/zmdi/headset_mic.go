package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetMic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 5q80 0 136 56.5T384 197v214q0 26-18.5 45T320 475H192v-43h149v-21h-85V240h85v-43q0-62-43.5-105.5T192 48T86.5 91.5T43 197v43h85v171H64q-27 0-45.5-19T0 347V197q0-79 56-135.5T192 5z"/>`),
		g.Group(children),
	)
}