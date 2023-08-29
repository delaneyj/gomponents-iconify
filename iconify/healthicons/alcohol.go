package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alcohol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 5a1 1 0 0 0-1-1H16a1 1 0 0 0-1 1v12.182a.715.715 0 0 1-.237.531A14.39 14.39 0 0 0 10 28.408v13.384c0 1.22.989 2.208 2.208 2.208h9.435A2.356 2.356 0 0 0 24 41.644V28.063c0-3.882-1.542-7.605-4.287-10.35a.727.727 0 0 1-.213-.514V5ZM12 29h10v8H12v-8Zm23.639-5H28.36l.64-6h6l.638 6ZM27.2 16h9.6l1.013 9.536A5.848 5.848 0 0 1 33 31.916V42h3a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2h3V31.915a5.848 5.848 0 0 1-4.813-6.379L27.2 16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}