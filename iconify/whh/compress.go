package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 192H800q-2 0 38 41t81 81t41 38v262q0 11-7.5 18.5T934 640h58q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 704H832q0 26-18.5 45T768 768H320v64h32q13 0 22.5 9.5T384 864v32h96q13 0 22.5 9.5T512 928t-9.5 22.5T480 960h-96v32q0 13-9.5 22.5T352 1024H160q-13 0-22.5-9.5T128 992v-32H32q-13 0-22.5-9.5T0 928t9.5-22.5T32 896h96v-32q0-13 9.5-22.5T160 832h32V64q0-27 18.5-45.5t45-18.5T301 18.5T320 64h448q27 0 45.5 18.5T832 128h160q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 192zm-390 0H320v448h282q-11 0-18.5-7.5T576 614V218q0-11 7.5-18.5T602 192zm166 64H640v320h256V384H768V256z"/>`),
		g.Group(children),
	)
}