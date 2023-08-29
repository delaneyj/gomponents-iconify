package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpwardsArrowWithLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 7h2.941v1.943H1z"/><path d="M14.02 7.049H8.017v1.909h6.052c.545 0 .988.44.988.979v3.129a.985.985 0 0 1-.988.979H7.958a.985.985 0 0 1-.988-.979V3.968h1.75c.392-.393.392-.638 0-1.029L6.708.291a1.003 1.003 0 0 0-1.42 0l-1.99 2.648c-.392.392-.392.638 0 1.029h1.729v9.053a2.977 2.977 0 0 0 2.98 2.967h6.014a2.976 2.976 0 0 0 2.98-2.967v-3.006a2.979 2.979 0 0 0-2.981-2.966z"/></g>`),
		g.Group(children),
	)
}