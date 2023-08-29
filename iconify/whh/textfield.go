package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textfield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 576H512v64h96q13 0 22.5 9.5T640 672t-9.5 22.5T608 704H352q-13 0-22.5-9.5T320 672t9.5-22.5T352 640h96v-64H64q-27 0-45.5-18.5T0 512V192q0-26 18.5-45T64 128h384V64h-96q-13 0-22.5-9.5T320 32t9.5-22.5T352 0h256q13 0 22.5 9.5T640 32t-9.5 22.5T608 64h-96v64h448q27 0 45.5 19t18.5 45v320q0 27-19 45.5T960 576zM448 192H64v320h384V192zm512 0H512v320h448V192zm-800 64q13 0 22.5 9.5T192 288v128q0 13-9.5 22.5T160 448t-22.5-9.5T128 416V288q0-13 9.5-22.5T160 256z"/>`),
		g.Group(children),
	)
}