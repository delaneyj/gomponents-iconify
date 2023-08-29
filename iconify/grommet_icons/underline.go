package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.41 4.53v-.35h4.66v.36h-.49a1.34 1.34 0 0 0-1.19.65a3 3 0 0 0-.2 1.4v5.33a9.45 9.45 0 0 1-.41 3.08a3.85 3.85 0 0 1-1.54 1.87a5.49 5.49 0 0 1-3.13.78a5.89 5.89 0 0 1-3.27-.75a4 4 0 0 1-1.58-2A11.14 11.14 0 0 1 7 11.64V6.5a2.58 2.58 0 0 0-.33-1.59a1.38 1.38 0 0 0-1.08-.38H5v-.35h5.68v.36h-.5A1.3 1.3 0 0 0 9.06 5a2.87 2.87 0 0 0-.25 1.5v5.73A12.52 12.52 0 0 0 9 14a3.71 3.71 0 0 0 .51 1.54a2.77 2.77 0 0 0 1.06.91a3.68 3.68 0 0 0 1.7.36a4.69 4.69 0 0 0 2.31-.56a3 3 0 0 0 1.39-1.44a8.33 8.33 0 0 0 .37-3V6.5A2.72 2.72 0 0 0 16 5a1.43 1.43 0 0 0-1.12-.43ZM4.93 20v-1H19v1Z"/>`),
		g.Group(children),
	)
}