package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DbTwoBufferPool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26.338 31l-1.736-1l2.288-4H23l3.993-7l1.737 1l-2.284 4h3.89l-3.998 7zM8 14h4v2H8zm12 0h4v2h-4zM8 18h4v2H8zm6 0h4v2h-4zm0 4h4v2h-4z"/><path fill="currentColor" d="M5 11h22v5h2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v22c0 1.103.897 2 2 2h15v-2H5V11Zm22-6v4H5V5h22Z"/>`),
		g.Group(children),
	)
}