package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M734 448h258q13 0 22.5 9.5t9.5 22.5t-9.5 22.5T992 512H672q-13 0-22.5-9.5T640 480q0-14 10-23L930 64H672q-13 0-22.5-9.5T640 32t9.5-22.5T672 0h320q12 0 22 11.5t10 20.5t-8 21zM576 800q0 13-9.5 22.5T544 832H288q-13 0-22.5-9.5T256 800q0-14 10-23l216-265H288q-13 0-22.5-9.5T256 480t9.5-22.5T288 448h256q12 0 22 11.5t10 20.5t-8 21L350 768h194q13 0 22.5 9.5T576 800zM192 992q0 13-9.5 22.5T160 1024H32q-13 0-22.5-9.5T0 992q0-14 10-23l88-137H32q-13 0-22.5-9.5T0 800t9.5-22.5T32 768h128q12 0 22 11.5t10 20.5t-8 21L94 960h66q13 0 22.5 9.5T192 992z"/>`),
		g.Group(children),
	)
}