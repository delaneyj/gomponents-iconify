package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IAmSober(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.36 13.347l5.166 5.166h0L24 27.038h0l-8.525-8.525h0l5.165-5.166a4.752 4.752 0 0 1 6.72 0Zm4.256 21.305l-5.166-5.165h0l8.525-8.525h0l8.526 8.525h0l-5.166 5.165a4.752 4.752 0 0 1-6.72 0Zm-21.951 0L4.5 29.487h0l8.525-8.526h0l8.525 8.526h0l-5.165 5.165a4.752 4.752 0 0 1-6.72 0Z"/>`),
		g.Group(children),
	)
}