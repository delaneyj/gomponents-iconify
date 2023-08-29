package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flagtriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 447L384 767q-27 0-45.5-18.5T320 703V64q0-27 18.5-45.5T384 0l576 319q2 1 5 3q14 6 21.5 10.5t17.5 12t15 17.5t5 21q0 13-5 23t-16 18t-19.5 12.5T960 447zM384 959.5q0 26.5-18.5 45T320 1023H64q-27 0-45.5-18.5T0 959.5T18.5 914T64 895h64V64q0-27 18.5-45.5T192 0t45.5 18.5T256 64v831h64q27 0 45.5 19t18.5 45.5z"/>`),
		g.Group(children),
	)
}