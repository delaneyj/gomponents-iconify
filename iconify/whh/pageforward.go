package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pageforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1006 427L587 755q-11 12-29.5 12.5t-32-7.5t-13.5-19V512H320q-26 0-45 18.5T256 576v384q0 26-18.5 45t-45.5 19H64q-26 0-45-19T0 960V512q0-106 75-181t181-75h256V28q0-12 13.5-20t32-7.5T587 13l419 327q18 18 18 43.5t-18 43.5z"/>`),
		g.Group(children),
	)
}