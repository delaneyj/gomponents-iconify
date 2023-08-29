package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RevolvingHearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRevolvingHearts0" d="M43 36.164a8.51 8.51 0 0 0-16-4.044a8.51 8.51 0 0 0-16 4.044a8.47 8.47 0 0 0 1.886 5.337l-.003.003L27 59l14.117-17.496l-.003-.002A8.47 8.47 0 0 0 43 36.163zm18-18.643a4.521 4.521 0 0 0-8.5-2.148a4.521 4.521 0 1 0-7.498 4.984H45l7.5 9.296l7.5-9.295l-.002-.001A4.5 4.5 0 0 0 61 17.52z"/></defs><use href="#openmojiRevolvingHearts0" fill="#FFA7C0"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path stroke-linecap="round" d="M19.614 21.692C23.718 18.496 29.541 16.5 36 16.5c.69 0 1.374.023 2.049.068M60.92 30.261c.38 1.202.58 2.453.58 3.739c0 8.728-9.241 15.873-20.94 16.46"/><use href="#openmojiRevolvingHearts0" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}