package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceBx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 5a3.003 3.003 0 0 0-3-3H4v13h5a3.003 3.003 0 0 0 3-3v-1.5a2.977 2.977 0 0 0-.78-2a2.977 2.977 0 0 0 .78-2zM6 4h3a1.001 1.001 0 0 1 1 1v1.551A.959.959 0 0 1 9 7.5H6zm4 8a1 1 0 0 1-1 1H6V9.5h3a1 1 0 0 1 1 1zm12-7h-2l-2 3.897L16 5h-2l2.905 5L14 15h2l2-3.799L20 15h2l-2.902-5L22 5z"/><circle cx="9" cy="27" r="1" fill="currentColor"/><path fill="currentColor" d="M2 18h4v2H2zm6 0h4v2H8zm6 0h4v2h-4zm6 0h4v2h-4zm6 0h4v2h-4zm0 13H6a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM6 25v4h20v-4z"/>`),
		g.Group(children),
	)
}