package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Removetags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.914 680l-325 325q-18 18-44.5 18t-44.5-18l-144-144v-62q0-40-28-68t-68-28h-62l-272-272q-18-18-18-79V64q0-27 19-45.5t45-18.5h288q62 0 80 18l574 573q18 19 18 45t-18 44zm-750-552q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5t-37.5-90.5t-90.5-37.5zm-192 639h256q26 0 45 19t19 45.5t-19 45t-45 18.5h-256q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19z"/>`),
		g.Group(children),
	)
}