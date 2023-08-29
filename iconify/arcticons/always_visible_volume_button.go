package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlwaysVisibleVolumeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.442 19.05c1.2.9 2.2 3.2 2 5.7c-.2 1.9-1.2 3.5-2 4.1m-18-9.1v8.7h5.8l7.5 6v-20.9l-7.5 6.2h-5.8Zm21.5-4.6c2.1 1.6 3.8 5.7 3.6 10.2c-.3 3.4-2 6.2-3.6 7.4"/>`),
		g.Group(children),
	)
}