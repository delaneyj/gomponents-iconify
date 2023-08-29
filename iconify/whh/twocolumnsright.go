package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twocolumnsright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.06 1024h-384q-53 0-90.5-37.5t-37.5-90.5V128q0-53 37.5-90.5T512.06 0h384q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-864q0-13-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v704q0 13 9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5V160zm-640 864h-192q-26 0-45-19t-19-45V64q0-26 19-45t45-19h192q26 0 45 19t19 45v896q0 26-18.5 45t-45.5 19z"/>`),
		g.Group(children),
	)
}