package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BootstrappingOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.667 12c-12 0-9 9-9 9v3h18v-3s3-9-9-9Zm7 10h-14v-1h14Zm0-2h-14c-.243-1.937-.117-6 7-6s7.243 4.064 7 6ZM8 6h8v2H8zm1-3h6v2H9z"/><path fill="currentColor" d="M12 0H3l1 13c1-1 2.033-2.5 8-2.5c5.97 0 7 1.5 8 2.5l1-13Zm6.308 9H5.692l-.538-7h13.692Z"/>`),
		g.Group(children),
	)
}