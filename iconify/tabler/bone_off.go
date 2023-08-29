package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.5 8.502l.38-.38A3 3 0 1 1 18 6a3 3 0 1 1-2.12 5.122l-.372.372M13.5 13.502l-2.378 2.378a3 3 0 1 1-5.117 2.297V18h-.176a3 3 0 1 1 2.298-5.115l2.378-2.378M3 3l18 18"/>`),
		g.Group(children),
	)
}