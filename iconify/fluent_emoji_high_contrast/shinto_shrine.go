package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShintoShrine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.993 2.416v-.02c-.03-.25-.25-.43-.49-.39c-8.69 1.22-18.31 1.22-27 0a.44.44 0 0 0-.49.39v.02c-.11.88.48 1.68 1.33 1.79c.22.03.44.05.66.07v1.14c0 .87.71 1.58 1.58 1.58H7v3H3.003c-.55 0-1 .45-1 1s.45 1 1 1H7v16h-.157c-.46 0-.84.38-.84.84v1.16h5v-1.16c0-.46-.38-.84-.84-.84H10v-13h.003v-3H22v16h-.077c-.46 0-.84.38-.84.84v1.16h5v-1.16c0-.46-.38-.84-.84-.84H25v-13h.003v-3h4c.56 0 1-.45 1-1s-.45-1-1-1H25v-3h1.433c.87 0 1.57-.71 1.57-1.58v-1.14c.22-.03.44-.05.66-.07c.85-.11 1.44-.91 1.33-1.79ZM22 9.996h-4v-3h4v3Zm-8 0h-4v-3h4v3Z"/>`),
		g.Group(children),
	)
}