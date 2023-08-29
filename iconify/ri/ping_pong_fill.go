package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PingPongFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 2a9.5 9.5 0 0 1 9.5 9.5c0 1.53-.361 2.974-1.003 4.254l2.463 2.464a1 1 0 0 1 0 1.414l-2.828 2.828a1 1 0 0 1-1.415 0l-2.463-2.463A9.462 9.462 0 0 1 11.5 21a9.5 9.5 0 0 1 0-19Zm5.303 13.388l-1.414 1.414l3.536 3.535l1.414-1.414l-3.536-3.535Zm1.864-6.105l-9.384 9.384A7.51 7.51 0 0 0 11.5 19a7.48 7.48 0 0 0 2.74-.516l-.972-.974a1 1 0 0 1 0-1.414l2.828-2.828a1 1 0 0 1 1.414 0l.974.972A7.48 7.48 0 0 0 19 11.5c0-.772-.117-1.516-.333-2.217Z"/>`),
		g.Group(children),
	)
}