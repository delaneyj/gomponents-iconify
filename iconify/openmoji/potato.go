package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Potato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiPotato0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M59.042 27.873c0 18.29-7.596 27.208-14.675 31.541c-6.185 3.804-11.992 4.107-11.992 4.107c-10.807 0-19.574-8.767-19.574-19.575c0-6.65 3.733-11.43 8.376-16.073C33.247 15.804 31.94 9.508 42.83 9.508c8.956 0 16.212 7.268 16.212 16.224v2.141z"/></defs><path fill="#A57939" d="M59.042 27.873c0 18.29-7.596 27.208-14.675 31.541c-6.185 3.804-11.992 4.107-11.992 4.107c-10.807 0-19.574-8.767-19.574-19.575c0-6.65 3.733-11.43 8.376-16.073C33.247 15.804 31.94 9.508 42.83 9.508c8.956 0 16.212 7.268 16.212 16.224v2.141z"/><path fill="#a57939" d="M59.045 25.73v2.14c0 18.29-7.6 27.21-14.68 31.54c-6.18 3.81-11.99 4.11-11.99 4.11c-1.67 0-3.28-.21-4.82-.6c31.56-8.62 25.35-49.89 25.35-49.89c3.74 2.97 6.14 7.56 6.14 12.7z"/><circle cx="38.274" cy="51.65" r="3.241" fill="#A57939"/><use href="#openmojiPotato0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="39.534" cy="16.215" r="1.26"/><circle cx="38.274" cy="51.65" r="3.241" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><use href="#openmojiPotato0" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}