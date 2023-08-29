package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOnceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M24 128a72.08 72.08 0 0 1 72-72h96V40a8 8 0 0 1 13.66-5.66l24 24a8 8 0 0 1 0 11.32l-24 24A8 8 0 0 1 192 88V72H96a56.06 56.06 0 0 0-56 56a8 8 0 0 1-16 0Zm200-8a8 8 0 0 0-8 8a56.06 56.06 0 0 1-56 56H64v-16a8 8 0 0 0-13.66-5.66l-24 24a8 8 0 0 0 0 11.32l24 24A8 8 0 0 0 64 216v-16h96a72.08 72.08 0 0 0 72-72a8 8 0 0 0-8-8Zm-88 40a8 8 0 0 0 8-8v-48a8 8 0 0 0-11.58-7.16l-16 8a8 8 0 1 0 7.16 14.31l4.42-2.21V152a8 8 0 0 0 8 8Z"/>`),
		g.Group(children),
	)
}