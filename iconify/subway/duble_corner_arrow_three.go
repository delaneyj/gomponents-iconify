package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M116.4 69.9H256L186.2.1L0 0l.1 186.2L69.9 256V116.4L209.5 256l46.5-46.5L116.4 69.9zm395.5 255.9L442.1 256v139.6L302.5 256L256 302.5l139.6 139.6H256l69.8 69.8l186.2.1l-.1-186.2z"/>`),
		g.Group(children),
	)
}