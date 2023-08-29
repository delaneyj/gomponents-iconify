package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M197.8 128L302.5 23.3L0 0l23.3 302.5L128 197.8l58.2 58.2l69.8-69.8l-58.2-58.2zM512 512l-23.3-302.5L384 314.2L325.8 256L256 325.8l58.2 58.2l-104.7 104.7L512 512z"/>`),
		g.Group(children),
	)
}