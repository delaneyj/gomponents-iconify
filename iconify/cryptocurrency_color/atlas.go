package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atlas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><path id="cryptocurrencyColorAtlas0" d="M2.5 4.938L0 0h5z"/><path id="cryptocurrencyColorAtlas1" d="M8.5 6.498L4.03 15.99c-.148.304-.225.55-.55.55l-2.953.002c-.423 0-.657-.109-.451-.55L7.296.447C7.445.19 7.537 0 7.862 0H9.14c.325 0 .417.19.565.448l7.22 15.544c.206.442-.028.551-.451.551l-2.953-.001c-.325 0-.402-.247-.55-.551L8.5 6.498z"/></defs><g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#31FAFB" fill-rule="nonzero"/><use fill="#FFF" href="#cryptocurrencyColorAtlas0" transform="translate(13.5 21.312)"/><use fill="#FFF" href="#cryptocurrencyColorAtlas1" transform="translate(7.5 6.25)"/></g>`),
		g.Group(children),
	)
}