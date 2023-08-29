package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlDeviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M17.57 3h-3.271l-.326.33l-1.544 2.942l-.486.327H6.432v4.495h3.03l.27.327l-3.3 6.305v3.273h3.272l.327-.33l1.543-2.943l.486-.326h5.511v-4.495h-3.03l-.269-.329l3.299-6.303L17.57 3z" fill="currentColor"/>`),
		g.Group(children),
	)
}