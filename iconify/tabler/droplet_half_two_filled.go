package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletHalfTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m13.905 2.923l.098.135l4.92 7.306a7.566 7.566 0 0 1 1.043 3.167l.024.326c.007.047.01.094.01.143l-.002.06c.056 2.3-.944 4.582-2.87 6.14c-2.969 2.402-7.286 2.402-10.255 0c-1.904-1.54-2.904-3.787-2.865-6.071a1.052 1.052 0 0 1 .013-.333a7.66 7.66 0 0 1 .913-3.176l.172-.302l4.893-7.26c.185-.275.426-.509.709-.686c1.055-.66 2.446-.413 3.197.55zm-2.06 1.107l-.077.038l-.041.03l-.037.036l-.033.042l-4.863 7.214A5.607 5.607 0 0 0 6.143 13h11.723a5.444 5.444 0 0 0-.49-1.313l-.141-.251l-4.891-7.261a.428.428 0 0 0-.5-.145z"/></g>`),
		g.Group(children),
	)
}