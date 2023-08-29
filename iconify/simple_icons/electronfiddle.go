package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electronfiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 0c-.631 0-1.143.512-1.143 1.143V15A4.573 4.573 0 0 0 8 24a4.571 4.571 0 0 0 1.143-8.999v-4.715h4.735c.54 0 .98-.512.98-1.143S14.417 8 13.877 8H9.143V2.286h10.286a1.143 1.143 0 1 0 0-2.286Zm0 17.143a2.286 2.286 0 1 1 0 4.571a2.286 2.286 0 0 1 0-4.571z"/>`),
		g.Group(children),
	)
}