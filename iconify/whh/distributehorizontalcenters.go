package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Distributehorizontalcenters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 768h-64v192q0 26-18.5 45t-45 19t-45.5-19t-19-45V768h-64q-27 0-45.5-19T512 704V320q0-27 18.5-45.5T576 256h64V64q0-26 19-45t45.5-19t45 19T768 64v192h64q27 0 45.5 18.5T896 320v384q0 27-18.5 45.5T832 768zm-64-352q0-13-9.5-22.5T736 384h-64q-13 0-22.5 9.5T640 416v192q0 13 9.5 22.5T672 640h64q13 0 22.5-9.5T768 608V416zM320 896h-64v64q0 26-18.5 45t-45 19t-45.5-19t-19-45v-64H64q-26 0-45-19T0 832V192q0-27 19-45.5T64 128h64V64q0-26 19-45t45.5-19t45 19T256 64v64h64q27 0 45.5 18.5T384 192v640q0 27-18.5 45.5T320 896z"/>`),
		g.Group(children),
	)
}