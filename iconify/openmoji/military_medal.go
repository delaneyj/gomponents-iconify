package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryMedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiMilitaryMedal0" d="m36 27.737l5.248 16.151H58.23L44.491 53.87l5.248 16.151L36 60.039l-13.739 9.982l5.248-16.151l-13.739-9.982h16.982zm0-8.673V2.021m20 6.958V2.021H16v6.958l20 14.758z"/></defs><g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#fcea2b" d="m36 27.737l5.248 16.151H58.23L44.491 53.87l5.248 16.151L36 60.039l-13.739 9.982l5.248-16.151l-13.739-9.982h16.982z"/><path fill="#ea5a47" d="M56 8.979V2.021H16v6.958l20 14.758z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiMilitaryMedal0"/><use href="#openmojiMilitaryMedal0"/></g>`),
		g.Group(children),
	)
}