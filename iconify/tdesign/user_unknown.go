package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserUnknown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM18 14c-.93 0-1.5.656-1.5 1.249v1h-2v-1C14.5 13.358 16.17 12 18 12s3.5 1.358 3.5 3.249a3.13 3.13 0 0 1-1.027 2.3L19 18.939v.683h-2v-1.546l2.112-1.993c.256-.235.388-.53.388-.834c0-.593-.57-1.249-1.5-1.249ZM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2H8Zm9 4.996h2.003V23h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}