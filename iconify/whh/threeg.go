package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Threeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.31 832h-384q-26 0-45-18.5t-19-45.5V64q0-27 19-45.5t45-18.5h384q27 0 45.5 18.5t18.5 45.5v64q0 27-18.5 45.5t-45.5 18.5h-256v448h128V512q-26 0-45-18.5t-19-45.5v-64q0-26 19-45t45-19h128q27 0 45.5 19t18.5 45v384q0 27-18.5 45.5t-45.5 18.5zm-576 0h-320q-27 0-45.5-18.5T.31 768v-64q0-26 18.5-45t45.5-19h192V512h-192q-26 0-45-18.5T.31 448v-64q0-26 19-45t45-19h192V192h-192q-27 0-45.5-18.5T.31 128V64q0-26 18.5-45t45.5-19h320q27 0 45.5 19t18.5 45v704q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}