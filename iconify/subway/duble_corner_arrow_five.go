package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209.5 23.3L314.2 128L256 186.2l69.8 69.8l58.2-58.2l104.7 104.7L512 0L209.5 23.3zM256 325.8L186.2 256L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384l58.2-58.2z"/>`),
		g.Group(children),
	)
}