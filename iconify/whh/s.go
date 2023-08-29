package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func S(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704.5 320q-26.5 0-45.5-18.5T640 256q0-53-37.5-90.5T512 128H257q-53 0-90.5 37.5T129 256v64q0 53 37.5 90.5T257 448h255q106 0 181 75t75 181v64q0 106-75 181t-181 75H256q-106 0-181-75T0 768q0-26 18.5-45t45-19t45.5 19t19 45q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768v-64q0-53-37.5-90.5T512 576H256q-106 0-180.5-75T1 320v-64Q1 150 76 75T257 0h255q106 0 181 75t75 181q0 27-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}