package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M341.3 0v85.3h-64c-70.7 0-128 57.3-128 128v85.3c0 35.4-28.6 64-64 64H0v64h85.3c70.7 0 128-57.3 128-128v-85.3c0-35.4 28.6-64 64-64h64v85.3L512 128v-21.3L341.3 0zM114 156.4l37.6-52.1c-19.4-11.8-42-19-66.3-19H0v64h85.3c10.4 0 20 2.7 28.7 7.1zm227.3 206.3h-64c-10.4 0-20-2.7-28.7-7.1L211 407.7c19.4 11.8 42 19 66.3 19h64V512L512 405.3V384L341.3 277.3v85.4z"/>`),
		g.Group(children),
	)
}