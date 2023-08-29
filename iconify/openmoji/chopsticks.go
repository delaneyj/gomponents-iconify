package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chopsticks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiChopsticks0" d="M53.76 58.958L18.83 20.3a2 2 0 0 0-2.828 2.828zm2.825-1.53l-26.073-45.11a2 2 0 0 0-3.357 2.174z"/></defs><use href="#openmojiChopsticks0" fill="#D22F27" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><use href="#openmojiChopsticks0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}