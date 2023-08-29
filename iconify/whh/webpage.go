package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webpage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.06 1024h-512q-27 0-45.5-19t-18.5-45V448q0-27 18.5-45.5t45.5-18.5h512q27 0 45.5 18.5t18.5 45.5v512q0 27-18.5 45.5t-45.5 18.5zm-64-480q0-13-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v320q0 13 9.5 22.5t22.5 9.5h320q13 0 22.5-9.5t9.5-22.5V544zm64-224h-512q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h512q27 0 45.5 19t18.5 45v192q0 27-18.5 45.5t-45.5 18.5zm-704 704h-192q-26 0-45-19t-19-45V64q0-26 19-45t45-19h192q27 0 45.5 19t18.5 45v896q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}