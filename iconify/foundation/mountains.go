package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m95.294 78.127l.026-.015L62.074 20.53l-.014.008c-.732-1.502-2.26-2.547-4.044-2.547a4.506 4.506 0 0 0-4.107 2.664l-22.125 38.32l-5.417-9.382l-.007.004a2.373 2.373 0 0 0-2.133-1.344c-.964 0-1.791.578-2.164 1.402L6.271 77.009H6.2v.122l-1.529 2.648a1.365 1.365 0 0 0-.285.832c0 .764.619 1.383 1.382 1.383c.152 0 .295-.03.432-.076v.091h87v-.05a2.613 2.613 0 0 0 2.413-2.601a2.588 2.588 0 0 0-.319-1.231z"/>`),
		g.Group(children),
	)
}