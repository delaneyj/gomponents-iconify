package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v20H2V2h20Zm0-2H2a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2Z"/><path fill="currentColor" d="M18 17H6v2h12v-2Zm-5 2h-2v3h2v-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M9.327 7.089A1.99 1.99 0 0 1 9.272 7H2V5h7.266A1.995 1.995 0 0 1 11 4h2a2.002 2.002 0 0 1 1.731 1H23v2h-8.272a1.997 1.997 0 0 1-1.186.916L12 11l-1.542-3.084a1.995 1.995 0 0 1-1.131-.827ZM12 7h-1.004h2.008H12Zm1-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}