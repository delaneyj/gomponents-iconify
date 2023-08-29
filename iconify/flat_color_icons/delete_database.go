package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteDatabase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2zm0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2zm0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		g.Group(children),
	)
}