package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiComment0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M23.437 50.356C22.105 53.88 18.432 58.712 15 61c7.165 0 12.233-2.092 16-7.273"/></defs><g fill="#FCEA2B"><circle cx="36" cy="35" r="20.8"/><path d="M23.437 50.356C22.105 53.88 18.432 58.712 15 61c7.165 0 12.233-2.092 16-7.273"/></g><circle cx="26" cy="35.079" r="2.857"/><circle cx="36" cy="35.079" r="2.857"/><circle cx="46" cy="35.079" r="2.857"/><use href="#openmojiComment0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><use href="#openmojiComment0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M20.48 47.617A19.917 19.917 0 0 1 16 35c0-11.046 8.954-20 20-20s20 8.954 20 20s-8.954 20-20 20c-.37 0-.738-.01-1.104-.03"/>`),
		g.Group(children),
	)
}