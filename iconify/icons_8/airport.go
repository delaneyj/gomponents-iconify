package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3c-1.645 0-3 1.355-3 3v7.344l-8.406 3.75l-.594.25v4.78L5.125 22L13 21.125v1.844l-2.563 1.717l-.437.28v4.253l1.188-.25L16 28l4.813.97l1.187.25v-4.25l-.438-.282L19 22.968v-1.843l7.875.875l1.125.125v-4.78l-.594-.25L19 13.344V6c0-1.645-1.355-3-3-3zm0 2c.565 0 1 .435 1 1v8.656l.594.25L26 18.656v1.22L18.125 19L17 18.875v5.187l.438.282L20 26.062v.72l-3.813-.75L16 25.97l-.188.06l-3.812.75v-.72l2.563-1.717l.437-.282v-5.186L13.875 19L6 19.875v-1.22l8.406-3.75l.594-.25V6c0-.565.435-1 1-1z"/>`),
		g.Group(children),
	)
}