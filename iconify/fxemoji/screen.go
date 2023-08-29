package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#00B1FF" d="M29.303 100.064h454.689c9.389 0 16.999 7.611 16.999 16.999v288.19c0 9.389-7.611 16.999-16.999 16.999H29.303c-9.389 0-16.999-7.611-16.999-16.999v-288.19c-.001-9.388 7.61-16.999 16.999-16.999z"/><path fill="#59CAFC" d="M29.303 100.064h454.689c9.389 0 16.999 7.611 16.999 16.999v288.19c0 9.389-7.611 16.999-16.999 16.999"/>`),
		g.Group(children),
	)
}