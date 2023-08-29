package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kibo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 204.546L204.44 0H512L0 512V204.546zM.236 512h509.968L316.933 316.201L.236 512z"/>`),
		g.Group(children),
	)
}