package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 1024H192q-21 0-51.5-34.5t-53.5-81T64 831q-1-24-16.5-81.5T16 618T0 480v-32q0-27 18.5-45.5T64 384q22 0 39 13.5t22 33.5q3-28 3-47q0-58 2-112.5t8-122T158 41t34-41q45 0 54.5 45.5T256 256q0 38 6.5 65t16 38.5t19 18T313 384h7q0-27 18.5-45.5T384 320q64 0 64 128q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v64q0-26 18.5-45t45-19t45.5 19t19 45v288q0 46-7 83t-18.5 59.5t-25 39.5t-27 24.5t-25 12t-18.5 5.5h-7z"/>`),
		g.Group(children),
	)
}