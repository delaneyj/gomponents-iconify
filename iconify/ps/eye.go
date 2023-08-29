package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M264 43Q102 43 10 181q-6 11 0 22q92 138 254 138t254-138q6-11 0-22Q426 43 264 43zm0 42q27 0 45.5 18.5T328 149q0 28-18.5 46T264 213t-45.5-18t-18.5-46q0-27 18.5-45.5T264 85zm0 214q-43 0-80.5-12t-63-31t-40-34T55 192q50-64 113-90q-11 28-11 47q0 45 31 76t76 31t76-31t31-76q0-19-11-47q70 29 113 90q-80 107-209 107z"/>`),
		g.Group(children),
	)
}