package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectExtent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.1 0v1.914H0v6h3V3h5.1V0h-6zm9 0v3h6V0h-6zm9 0v3h6V0h-6zm9 0v3h6V0h-6zm9 0v3h6V0h-6zm9 0v3h6V0h-6zm9 0v3h6V0h-6zm9 0v3h1.8v1.2h3V0h-4.8zm1.8 7.2v6h3v-6h-3zM0 10.913v6h3v-6H0zM66.9 16.2v6h3v-6h-3zM0 19.914v6h3v-6H0zM66.9 25.2v6h3v-6h-3zM0 28.914v6h3v-6H0zM66.9 34.2v6h3v-6h-3zM0 37.914v6h3v-6H0zM66.9 43.2v6h3v-6h-3zM0 46.914v6h3v-6H0zM66.9 52.2v6h3v-6h-3zM0 55.914v5.191h3.809v-3H3v-2.19H0zm6.809 2.191v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9.648 1.899a2.076 2.076 0 0 0-2.19 2.324l3.137 33.676c.2 1.635 2.135 2.399 3.397 1.34l6.623-5.371l2.969 5.142c1.707 2.958 4.417 3.684 7.375 1.977c2.957-1.708 3.684-4.417 1.976-7.375l-2.959-5.125l7.848-3.008c1.548-.564 1.855-2.62.539-3.611L71.576 60.416a2.073 2.073 0 0 0-1.119-.412z" color="currentColor"/>`),
		g.Group(children),
	)
}