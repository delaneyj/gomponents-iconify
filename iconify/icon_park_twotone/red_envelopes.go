package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedEnvelopes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRedEnvelopes0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M40 4H8v40h32V4Z"/><path stroke-linecap="round" d="M40 10c-4 5.762-7.68 9.32-11.68 10.892m-8.583-.114C15.737 19.206 12 15.762 8 10l11.737 10.777Z"/><path fill="#555" d="M24 27a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRedEnvelopes0)"/>`),
		g.Group(children),
	)
}