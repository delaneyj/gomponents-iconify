package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seafile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.23 12.44a5.85 5.85 0 0 1 5.95 4.68c1.7-2.81 10.79-1.74 10.74 4.77c4-1.63 7.66 2.37 7.66 5.47s-1.7 5.27-5.28 5.27h-6c-3.58 0-9.74-1.3-10.3-5.1a4 4 0 0 1 2.45-3.25A1.93 1.93 0 0 0 26 26.71a2.44 2.44 0 0 0 2.74-2c.2-1.69-1.23-2.94-3.55-2.94a9 9 0 0 0-7 2.53c-2.59 2.36-2.38 8.35-9.24 8.35s-5.26-12.3 1.9-10.74c-2.74-4.93 3.25-9.56 8.38-9.45Z"/>`),
		g.Group(children),
	)
}