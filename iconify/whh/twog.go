package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.585 832h-384q-27 0-45.5-18.5t-18.5-45.5V64q0-27 18.5-45.5t45.5-18.5h384q27 0 45.5 18.5t18.5 45.5v64q0 26-18.5 45t-45.5 19h-256v448h128V512q-27 0-45.5-18.5t-18.5-45.5v-64q0-26 18.5-45t45.5-19h128q26 0 45 19t19 45v384q0 27-18.5 45.5t-45.5 18.5zm-576 0h-320q-27 0-45.5-18.5T.585 768V384q0-27 18.5-45.5t45.5-18.5h192V192h-192q-27 0-45.5-18.5T.585 128V64q0-27 18.5-45.5T64.585 0h320q27 0 45.5 18.5t18.5 45.5v384q0 26-18.5 45t-45.5 19h-192v128h192q27 0 45.5 18.5t18.5 45.5v64q0 26-18.5 45t-45.5 19z"/>`),
		g.Group(children),
	)
}