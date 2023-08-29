package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M256 1536h768v-512H256v512zm1024-512h512V256h-768v256h96q66 0 113 47t47 113v352zm768-864v960q0 66-47 113t-113 47h-608v352q0 66-47 113t-113 47H160q-66 0-113-47T0 1632V672q0-66 47-113t113-47h608V160q0-66 47-113T928 0h960q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}