package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pictureframe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 1024H32q-13 0-22.5-9.5T0 992V288q0-13 9.5-22.5T32 256h211L488 11l.5-1l.5-1q10-9 23-9t23 9l.5 1l.5 1l245 245h211q13 0 22.5 9.5t9.5 22.5v704q0 13-9.5 22.5T992 1024zM512 77L333 256h358zm384 307H128v512h768V384z"/>`),
		g.Group(children),
	)
}