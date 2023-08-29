package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 107q-62 0-105.5 43.5T107 256t43.5 105.5T256 405t105.5-43.5T405 256t-43.5-105.5T256 107zm0 256q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31zm21-299V21q0-21-21-21t-21 21v43q0 21 21 21t21-21zM85 256q0-21-21-21H21q-21 0-21 21t21 21h43q21 0 21-21zm406-21h-43q-21 0-21 21t21 21h43q21 0 21-21t-21-21zM256 427q-21 0-21 21v43q0 21 21 21t21-21v-43q0-21-21-21zm164-305l22-22q14-14 0-30q-16-14-30 0l-22 22q-14 14 0 30q8 6 15 6q9 0 15-6zm-313 6q7 0 15-6q14-16 0-30l-22-22q-14-14-30 0q-14 16 0 30l22 22q6 6 15 6zm313 262q-14-14-30 0q-14 16 0 30l22 22q6 6 15 6q7 0 15-6q14-16 0-30zm-328 0l-22 22q-14 14 0 30q8 6 15 6q9 0 15-6l22-22q14-14 0-30q-16-14-30 0z"/>`),
		g.Group(children),
	)
}