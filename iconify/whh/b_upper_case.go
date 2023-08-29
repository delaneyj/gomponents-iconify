package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 704v64q0 106-75 181t-181 75H64q-27 0-45.5-18.5T0 960V64q0-26 18.5-45T64 0h448q106 0 181 75t75 181v64q0 56-23.5 106T681 512q40 36 63.5 86T768 704zM640 256q0-53-37.5-90.5T512 128H160q-13 0-22.5 9.5T128 160v256q0 13 9.5 22.5T160 448h352q53 0 90.5-37.5T640 320v-64zm0 448q0-53-37.5-90.5T512 576H160q-13 0-22.5 9.5T128 608v256q0 13 9.5 22.5T160 896h352q53 0 90.5-37.5T640 768v-64z"/>`),
		g.Group(children),
	)
}