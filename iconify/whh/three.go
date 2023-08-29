package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Three(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 704v64q0 106-75 181t-181 75H256q-106 0-181-75T0 768q0-26 18.5-45t45-19t45.5 19t19 45q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768v-64q0-53-37.5-90.5T512 576h-64q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h64q53 0 90.5-37.5T640 320v-64q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256q0 27-19 45.5T63.5 320t-45-18.5T0 256Q0 150 75 75T256 0h256q106 0 181 75t75 181v64q0 56-23.5 106T681 512q40 36 63.5 86T768 704z"/>`),
		g.Group(children),
	)
}