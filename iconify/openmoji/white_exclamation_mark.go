package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiWhiteExclamationMark0" d="M35.987 49.373a2.5 2.5 0 0 1-2.5-2.5v-35.21a2.5 2.5 0 1 1 5 0v35.21a2.5 2.5 0 0 1-2.5 2.5z"/></defs><g fill="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36.093" cy="58.842" r="3"/><use href="#openmojiWhiteExclamationMark0"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36.093" cy="58.842" r="3"/><use href="#openmojiWhiteExclamationMark0"/></g>`),
		g.Group(children),
	)
}