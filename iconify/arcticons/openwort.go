package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openwort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.635 20.039L24 5.5l-9.25 37l-9.25-37m28.907 32.373L33.25 42.5l-1.157-4.627M42.5 5.5l-3.635 14.539M22.573 31.161h19.83v4.412h-19.83z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.573 26.75h19.83v4.412h-19.83zm0-4.412h19.83v4.412h-19.83z"/>`),
		g.Group(children),
	)
}