package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.072 20.3a2.999 2.999 0 0 0 3.856 0a3.002 3.002 0 0 0 .67-3.798l-2.095-3.227a.6.6 0 0 0-1.005 0L3.4 16.502a3.003 3.003 0 0 0 .671 3.798zm12 0a2.999 2.999 0 0 0 3.856 0a3.002 3.002 0 0 0 .67-3.798l-2.095-3.227a.6.6 0 0 0-1.005 0L15.4 16.502a3.003 3.003 0 0 0 .671 3.798zm-6-10a2.999 2.999 0 0 0 3.856 0a3.002 3.002 0 0 0 .67-3.798l-2.095-3.227a.6.6 0 0 0-1.005 0L9.4 6.502a3.003 3.003 0 0 0 .671 3.798z"/>`),
		g.Group(children),
	)
}