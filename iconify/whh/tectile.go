package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tectile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640H512q-53 0-90.5 37.5T384 768v192q0 26-18.5 45t-45.5 19H64q-26 0-45-19T0 960V704q0-26 19-45t45-19h192q53 0 90.5-37.5T384 512V64q0-26 19-45t45-19h512q27 0 45.5 19t18.5 45v512q0 27-19 45.5T960 640z"/>`),
		g.Group(children),
	)
}