package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomePercentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.53 11.03l-1.06-1.06l-6 6l1.06 1.06l6-6m-4.9-.66c-.23-.24-.55-.37-.88-.37c-.33 0-.65.13-.88.37c-.24.23-.37.55-.37.88c0 .33.13.65.37.88c.23.24.55.37.88.37c.33 0 .65-.13.88-.37c.24-.23.37-.55.37-.88c0-.33-.13-.65-.37-.88m4.5 4.5c-.23-.24-.55-.37-.88-.37c-.33 0-.65.13-.88.37c-.24.23-.37.55-.37.88c0 .33.13.65.37.88c.23.24.55.37.88.37c.33 0 .65-.13.88-.37c.24-.23.37-.55.37-.88c0-.33-.13-.65-.37-.88M19 12v8H5v-8H2l10-9l10 9h-3m-2-1.8l-5-4.5l-5 4.5V18h10v-7.8Z"/>`),
		g.Group(children),
	)
}