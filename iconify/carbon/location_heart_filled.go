package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHeartFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M18.303 9a2.662 2.662 0 0 0-1.908.806l-.393.405l-.397-.405a2.662 2.662 0 0 0-3.816 0a2.8 2.8 0 0 0 0 3.896L16.002 18l4.209-4.298a2.8 2.8 0 0 0 0-3.896A2.662 2.662 0 0 0 18.303 9Z"/><path fill="currentColor" d="M16 2A11.013 11.013 0 0 0 5 13a10.889 10.889 0 0 0 2.216 6.6s.3.395.349.452L16 30l8.439-9.953c.044-.053.345-.447.345-.447l.001-.003A10.885 10.885 0 0 0 27 13A11.013 11.013 0 0 0 16 2Zm4.21 11.702L16.003 18l-4.213-4.298a2.8 2.8 0 0 1 0-3.896a2.662 2.662 0 0 1 3.816 0l.397.405l.393-.405a2.662 2.662 0 0 1 3.816 0a2.8 2.8 0 0 1 0 3.896Z"/>`),
		g.Group(children),
	)
}