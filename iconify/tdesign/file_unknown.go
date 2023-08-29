package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUnknown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h9.5v2H3V1Zm12 2.414V7h3.586L15 3.414ZM18 15c-.93 0-1.5.656-1.5 1.249v1h-2v-1C14.5 14.358 16.17 13 18 13s3.5 1.358 3.5 3.249a3.13 3.13 0 0 1-1.027 2.3L19 19.939v.683h-2v-1.546l2.112-1.993c.256-.235.388-.53.388-.834c0-.593-.57-1.249-1.5-1.249Zm-1 6.996h2.003V24h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}