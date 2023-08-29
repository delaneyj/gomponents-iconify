package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nebulo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.27" cy="26.77" r="3.59" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.86 26.77h9.8m-3.28 2.85v-2.85m-10.11-9.09h-6.22a9.56 9.56 0 0 0 0 19.11h21.2a8.25 8.25 0 0 0 0-16.5h-.2a10.81 10.81 0 0 0-20.57-2.6"/>`),
		g.Group(children),
	)
}