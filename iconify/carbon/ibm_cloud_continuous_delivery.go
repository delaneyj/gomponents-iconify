package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudContinuousDelivery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 1l-6 6l1.415 1.402L22 4.818V21H10V10H8v11c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V4.815l3.586 3.587L29 7l-6-6Z"/><path fill="currentColor" d="M18.5 19h-5c-.827 0-1.5-.673-1.5-1.5v-5c0-.827.673-1.5 1.5-1.5h5c.827 0 1.5.673 1.5 1.5v5c0 .827-.673 1.5-1.5 1.5zM14 17h4v-4h-4v4zm2 14v-2c7.168 0 13-5.832 13-13c0-1.265-.181-2.514-.538-3.715l1.917-.57C30.79 13.1 31 14.542 31 16c0 8.271-6.729 15-15 15zM1.621 20.285A15.011 15.011 0 0 1 1 16C1 7.729 7.729 1 16 1v2C8.832 3 3 8.832 3 16c0 1.265.181 2.514.538 3.715l-1.917.57z"/>`),
		g.Group(children),
	)
}