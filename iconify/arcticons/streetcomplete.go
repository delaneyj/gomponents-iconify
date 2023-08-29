package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Streetcomplete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.28 5.5h5.44v4.59h4.16v9.17h-4.16v3.3h11l4.59 4.59l-4.59 4.59h-11V44.5h-5.44V31.74h-4.16v-9.18h4.16v-3.3h-11l-4.63-4.58l4.59-4.59h11Zm0 4.59h5.44m0 9.17h-5.44m0 3.3h5.44m.25 9.18h-5.69"/>`),
		g.Group(children),
	)
}