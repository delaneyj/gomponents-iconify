package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 448H890q-23 137-129.5 228.5T512 768q-78 0-148-29.5T240 656l91-91q75 75 181 75t181-75t75-181v-32q0-13 9.5-22.5T800 320h192q13 0 22.5 9.5t9.5 22.5v64q0 13-9.5 22.5T992 448zm-352-64q0 53-37.5 90.5T512 512t-90.5-37.5T384 384t37.5-90.5T512 256t90.5 37.5T640 384zM512 128q-106 0-181 75t-75 181v32q0 13-9.5 22.5T224 448H32q-13 0-22.5-9.5T0 416v-64q0-13 9.5-22.5T32 320h102q23-137 129.5-228.5T512 0q78 0 148 29.5T784 112l-91 91q-75-75-181-75z"/>`),
		g.Group(children),
	)
}