package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fourimages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-64V832h192v64q0 53-37.5 90.5t-90.5 37.5zm-64-1024h64q53 0 90.5 37.5t37.5 90.5v512h-192V0zm-572 623q-4 7-5 19.5t0 20.5l1 9v128q0 13 9.5 22.5t22.5 9.5h352v192h-512q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h479zm206 17l174-311v311h-174z"/>`),
		g.Group(children),
	)
}