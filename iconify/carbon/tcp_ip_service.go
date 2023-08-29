package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TcpIpService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28h14v2H16zm11-2h-8c-1.103 0-2-.897-2-2v-5c0-1.103.897-2 2-2h8c1.103 0 2 .897 2 2v5c0 1.103-.897 2-2 2zm-8-7v5h8v-5h-8zm-4 4h-5c-1.103 0-2-.897-2-2v-4h2v4h5v2zM2 13h14v2H2zm11-2H5c-1.103 0-2-.897-2-2V4c0-1.103.897-2 2-2h8c1.103 0 2 .897 2 2v5c0 1.103-.897 2-2 2zM5 4v5h8V4H5z"/>`),
		g.Group(children),
	)
}