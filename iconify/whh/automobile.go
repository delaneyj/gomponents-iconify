package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 750v82q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5v-64H192v64q0 27-18.5 45.5t-45 18.5T83 877.5T64 832v-82q-29-17-46.5-46T0 640V448q0-28 23-64.5T81 314l47-186q0-53 37.5-90.5T256 0h512q53 0 90.5 37.5T896 128l47 186q35 33 58 69.5t23 64.5v192q0 35-17.5 64T960 750zM128 576.5q0 26.5 18.5 45t45 18.5t45.5-18.5t19-45t-19-45.5t-45.5-19t-45 19t-18.5 45.5zM800 192q0-26-18.5-45T736 128H288q-26 0-45 19t-19 45l-32 128q0 27 18.5 45.5T256 384h512q27 0 45.5-18.5T832 320zm32.5 320q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19z"/>`),
		g.Group(children),
	)
}