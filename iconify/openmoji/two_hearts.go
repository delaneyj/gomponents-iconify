package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoHearts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiTwoHearts0" d="M43 36.164a8.51 8.51 0 0 0-16-4.044a8.51 8.51 0 0 0-16 4.044a8.47 8.47 0 0 0 1.886 5.337l-.003.003L27 59l14.117-17.496l-.003-.002A8.47 8.47 0 0 0 43 36.163zm18-18.643a4.521 4.521 0 0 0-8.5-2.148a4.521 4.521 0 1 0-7.498 4.984H45l7.5 9.296l7.5-9.295l-.002-.001A4.5 4.5 0 0 0 61 17.52z"/></defs><use href="#openmojiTwoHearts0" fill="#FFA7C0"/><use href="#openmojiTwoHearts0" fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}