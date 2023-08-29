package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiThreeOclock0" d="M36 18.989v17m11.63.004H36.083"/></defs><g fill="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="35.958" cy="35.99" r="23"/><use href="#openmojiThreeOclock0"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="35.958" cy="35.99" r="23"/><use href="#openmojiThreeOclock0"/></g>`),
		g.Group(children),
	)
}