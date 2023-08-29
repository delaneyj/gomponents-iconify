package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.5 0-4.25-1.75T6 15q0-2.25 1.438-3.913T11 9.075V3h2v6.075q2.15.35 3.575 2.013T18 15q0 2.5-1.75 4.25T12 21Zm0-2q1.4 0 2.463-.85T15.875 16h-7.75q.35 1.3 1.413 2.15T12 19Zm-3.875-5h7.75q-.35-1.3-1.413-2.15T12 11q-1.4 0-2.463.85T8.125 14Z"/>`),
		g.Group(children),
	)
}