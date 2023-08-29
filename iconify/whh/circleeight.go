package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-448q0-57-43-96q43-38 43-96q0-53-37.5-90.5T576 256H448q-53 0-90.5 37.5T320 384q0 58 43 96q-43 39-43 96v64q0 53 37.5 90.5T448 768h128q53 0 90.5-37.5T704 640v-64zM576 704H448q-27 0-45.5-18.5T384 640v-64q0-26 18.5-45t45.5-19h128q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5T576 704zm0-256H448q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128q27 0 45.5 19t18.5 45.5t-18.5 45T576 448z"/>`),
		g.Group(children),
	)
}