package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opentodolist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.87 37.75a16.77 16.77 0 1 1 0-27.5a.36.36 0 0 1 .13.49a.35.35 0 0 1-.48.13a14.15 14.15 0 1 0 0 26.26a.35.35 0 0 1 .45.25a.37.37 0 0 1-.1.37Zm-2.06-7.38c-1.48 2.19-2.18 4.82-6.1 4.82c-4.14 0-5.44-9.25-5.38-10.56s4.29-2.41 4.8-2.29s1.63 3.89 2.29 3.88s3.55-4.58 5.42-7s7.09-6.79 8.51-7.27a9.4 9.4 0 0 1 5.07.29c1.2.44-9.2 10.12-14.61 18.13Z"/>`),
		g.Group(children),
	)
}