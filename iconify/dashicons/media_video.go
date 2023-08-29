package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zm0 4h3l-3-3v3zm-1 8v-3c0-.27-.1-.51-.29-.71c-.2-.19-.44-.29-.71-.29H7c-.27 0-.51.1-.71.29c-.19.2-.29.44-.29.71v3c0 .27.1.51.29.71c.2.19.44.29.71.29h3c.27 0 .51-.1.71-.29c.19-.2.29-.44.29-.71zm3 1v-5l-2 2v1z"/>`),
		g.Group(children),
	)
}