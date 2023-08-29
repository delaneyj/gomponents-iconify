package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 10v11h9v2h7v-2h16V10H0zm1.777 1.777h7.114v7.668H7.11v-5.888H5.334v5.888H1.777v-7.668zm8.891 0h7.111v7.666h-3.556v1.78h-3.555v-9.446zm8.889 0h10.668v7.668h-1.78v-5.888h-1.777v5.888h-1.777v-5.888h-1.78v5.888h-3.554v-7.668zm-5.334 1.78v4.111H16v-4.111h-1.777z"/>`),
		g.Group(children),
	)
}