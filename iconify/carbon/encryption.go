package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Encryption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 21.278V19a4 4 0 0 0-8 0v2.278A1.994 1.994 0 0 0 20 23v5a2.002 2.002 0 0 0 2 2h6a2.002 2.002 0 0 0 2-2v-5a1.994 1.994 0 0 0-1-1.722zM25 17a2.002 2.002 0 0 1 2 2v2h-4v-2a2.002 2.002 0 0 1 2-2zm-3 11v-5h6v5zM2 2h2v4H2zm12 0h2v4h-2zm4 0h2v4h-2zM2 8h2v8H2zm0 10h2v8H2zm12 0h2v8h-2zM6 8h2v8H6zm12 0h2v6h-2zm-8 18H8a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-2-6v4h2v-4zm6-4h-2a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-2-6v4h2v-4zm-2-4H8a2.002 2.002 0 0 1-2-2V2h2v2h2V2h2v2a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}