package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 704v64q0 106-75 181t-181 75H64q-27 0-45.5-19T0 960V64q0-26 18.5-45T64 0h448q106 0 181 75t75 181v32q0 56-23.5 106T681 480q38 34 62.5 99T768 704zM480 192H224q-13 0-22.5 9.5T192 224v128q0 13 9.5 22.5T224 384h256q39 0 67.5-28t28.5-68t-28.5-68t-67.5-28zm-32 384H224q-13 0-22.5 9.5T192 608v192q0 13 9.5 22.5T224 832h224q53 0 90.5-37.5T576 704t-37.5-90.5T448 576z"/>`),
		g.Group(children),
	)
}