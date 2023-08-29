package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M505 122V85h109V15H243l-59 45v319s82 61 118 136c36 76 34 164 34 164h94s4-72 0-129c-5-58-24-129-24-129h260v-82H505v-37h161v-73H505v-37h130v-70H505zM0 64h142v308H0V64z"/>`),
		g.Group(children),
	)
}