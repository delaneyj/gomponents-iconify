package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kidney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiKidney0" d="M25.942 31.266a11.86 11.86 0 0 1-3.3-5.997c-.446-2.251-.284-4.934 2.187-6.594c0 0 10.439-7.388 19.593 0c0 0 11.08 6.825 9.394 24.65c0 0-.16 12.286-12.687 12.447c0 0-6.182 1.525-6.825-8.592c.061-.942.06-1.888-.005-2.83a14.336 14.336 0 0 0-4.977-9.96a59.361 59.361 0 0 1-3.38-3.124z"/></defs><g fill="#D22F27"><use href="#openmojiKidney0"/><use href="#openmojiKidney0"/></g><g fill="none" stroke="#000" stroke-width="2"><use href="#openmojiKidney0" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M34.987 29.073s-3.904 1.592-7.252.435m12.672 13.978s-6.906-.803-9.636-3.814c0 0-3.309-3.739-9.836 0l-6.778 2.337"/><path stroke-linecap="round" stroke-linejoin="round" d="M36.788 33.69a2.59 2.59 0 0 0-.844 1.44a3.076 3.076 0 0 0 .216 1.914c.338.786.785 2.584-.304 5.44m-21.699-.475l6.778-2.337s5.62-2.248 6.424-5.7c0 0-.322-6.023 2.087-7.308"/></g>`),
		g.Group(children),
	)
}