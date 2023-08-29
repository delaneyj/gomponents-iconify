package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 448H576V320h416q13 0 22.5 9t9.5 23v64q0 13-9.5 22.5T992 448zM640 256H320q-53 0-90.5-37.5T192 128t37.5-90.5T320 0q54 0 80.5 13.5t33.5 40t9.5 56T462 180t50 76v-43.5l1-40.5l3.5-40l6.5-35.5l10.5-33l16-25.5L572 17l30-12l38-5q53 0 90.5 37.5T768 128t-37.5 90.5T640 256zM319.5 64q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45 19t45.5-19t19-45.5t-19-45T319.5 64zm320 0q-26.5 0-45 18.5t-18.5 45t19 45.5t45 19t45-19t19-45.5t-19-45T639.5 64zM448 448H32q-13 0-22.5-9.5T0 416v-64q0-14 9.5-23t22.5-9h416v128zm0 576H128q-27 0-45.5-19T64 960V512h384v512zm512-512v448q0 26-18.5 45t-45.5 19H576V512h384z"/>`),
		g.Group(children),
	)
}