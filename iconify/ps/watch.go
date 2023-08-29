package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 256q0 39 20 73.5T75 384h-8v85q0 18 12.5 30.5T109 512h86q17 0 29.5-12.5T237 469v-85h-8q61-37 70-107h2q22 0 22-21t-22-21h-2q-9-68-70-107h8V43q0-18-12.5-30.5T195 0h-86Q92 0 79.5 12.5T67 43v85h8q-32 20-52 54.5T3 256zm106 213v-21h86v21h-86zm86-42h-86v-28q20 6 43 6q24 0 43-6v28zm64-171q0 45-31 76t-76 31t-76-31t-31-76t31-76t76-31t76 31t31 76zM109 43h86v21h-86V43zm0 42h86v28q-19-6-43-6q-23 0-43 6V85zm64 107h-42v73l49 49l30-30l-37-37v-55z"/>`),
		g.Group(children),
	)
}