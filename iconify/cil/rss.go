package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M66 256v32c88.225 0 160 71.776 160 160h32c0-105.869-86.131-192-192-192Z"/><path fill="currentColor" d="M66 140v32c152.187 0 276 123.813 276 276h32a305.982 305.982 0 0 0-90.211-217.789A305.987 305.987 0 0 0 66 140Z"/><path fill="currentColor" d="M456.674 282.955a422.588 422.588 0 0 0-90.861-134.768A422.724 422.724 0 0 0 66 24v32c216.149 0 392 175.851 392 392h32a421.378 421.378 0 0 0-33.326-165.045ZM90 360a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64Zm0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}