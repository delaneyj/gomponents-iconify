package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librephotosmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.864 39.114l14.25-14.25a1.222 1.222 0 0 0 0-1.728h0l-14.25-14.25a1.222 1.222 0 0 0-1.728 0h0l-14.25 14.25a1.222 1.222 0 0 0 0 1.728h0l14.25 14.25a1.222 1.222 0 0 0 1.728 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.864 43.142l14.25-14.25a1.222 1.222 0 0 0 0-1.728h0l-14.25-14.25a1.222 1.222 0 0 0-1.728 0h0l-14.25 14.25a1.222 1.222 0 0 0 0 1.728h0l14.25 14.25a1.222 1.222 0 0 0 1.728 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.864 35.086l14.25-14.25a1.222 1.222 0 0 0 0-1.728h0l-14.25-14.25a1.222 1.222 0 0 0-1.728 0h0l-14.25 14.25a1.222 1.222 0 0 0 0 1.728h0l14.25 14.25a1.222 1.222 0 0 0 1.728 0Z"/>`),
		g.Group(children),
	)
}