package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H64q-26 0-45-19T0 960V64q0-26 18.5-45T64 0h640q27 0 45.5 18.5T768 64v896q0 27-19 45.5t-45 18.5zM384 64q-53 0-90.5 37.5T256 192t37.5 90.5T384 320t90.5-37.5T512 192t-37.5-90.5T384 64zm0 384q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75zm0 384q-53 0-90.5-37.5T256 704t37.5-90.5T384 576t90.5 37.5T512 704t-37.5 90.5T384 832zm0-576q-26 0-45-18.5t-19-45t19-45.5t45.5-19t45 19t18.5 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}