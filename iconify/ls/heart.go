package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M582 30c166 0 221 192 163 297c-89 163-361 336-361 336S111 490 23 327C-35 222 19 30 186 30c154 0 193 133 198 153c5-20 44-153 198-153z"/>`),
		g.Group(children),
	)
}