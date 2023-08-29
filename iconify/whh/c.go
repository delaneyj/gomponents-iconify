package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func C(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704.5 320q-26.5 0-45-18.5T641 256q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256v512q0 53 37.5 90.5T256 896h256q54 0 91.5-37.5T641 768q0-27 18.5-45.5t45-18.5t45.5 19t19 45q0 106-75 181t-182 75H256q-106 0-181-75T0 768V256Q0 150 75 75T256 0h256q107 0 182 75t75 181q0 27-19 45.5T704.5 320z"/>`),
		g.Group(children),
	)
}