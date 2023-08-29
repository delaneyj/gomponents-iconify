package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tetristhree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 1024H64q-26 0-45-19T0 960V64q0-26 18.5-45T64 0h256q27 0 45.5 19T384 64v576h256q27 0 45.5 19t18.5 45v256q0 27-18.5 45.5T640 1024zM320 96q0-13-9.5-22.5T288 64H96q-13 0-22.5 9.5T64 96v192q0 13 9.5 22.5T96 320h192q13 0 22.5-9.5T320 288V96zm0 314q0-11-7.5-18.5T294 384H90q-11 0-18.5 7.5T64 410v204q0 11 7.5 18.5T90 640h204q11 0 18.5-7.5T320 614V410zm0 320q0-11-7.5-18.5T294 704H90q-11 0-18.5 7.5T64 730v204q0 11 7.5 18.5T90 960h204q11 0 18.5-7.5T320 934V730zm320 0q0-11-7.5-18.5T614 704H410q-11 0-18.5 7.5T384 730v204q0 11 7.5 18.5T410 960h204q11 0 18.5-7.5T640 934V730z"/>`),
		g.Group(children),
	)
}