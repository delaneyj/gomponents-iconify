package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleMessageHi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.979.021c-4.398 0-7.968 2.792-7.968 6.235c0 3.143 2.976 5.734 6.841 6.164l-2.047 3.551l6.119-3.926c2.94-.917 5.023-3.161 5.023-5.789c0-3.444-3.568-6.235-7.968-6.235zm.061 8.997H7.972V7.075h-1.94v1.943l-1.072.023V3.992l1.072-.023v2.018h1.94V3.969H9.04v5.049zm1.96 0H9.954v-3.08H11v3.08zm.031-3.977H9.954V3.937h1.077v1.104zm3 3.975h-1.063V7.963h1.063v1.053zm-.015-2h-1.047V3.954h1.047v3.062z"/>`),
		g.Group(children),
	)
}