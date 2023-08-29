package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 672q0 93-65.5 158.5T480 896h-32v64q0 27-18.5 45.5T384 1024t-45.5-18.5T320 960v-64h-64v64q0 27-18.5 45.5T192 1024t-45.5-18.5T128 960v-64H64q-27 0-45.5-18.5T0 832t18.5-45.5T64 768h64V256H64q-27 0-45.5-18.5T0 192t18.5-45.5T64 128h64V64q0-26 18.5-45T192 0t45.5 19T256 64v64h64V64q0-26 18.5-45t45-19T429 19t19 45v67q82 11 137 74t55 147q0 70-42 130q48 30 77 80.5T704 672zM416 256H256v192h160q40 0 68-28t28-68t-28-68t-68-28zm64 320H256v192h224q40 0 68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}