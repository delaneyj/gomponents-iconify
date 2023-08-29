package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAddressBook0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 36v8h32V4H8v8M6 30h4m-4-6h4m-4-6h4"/><circle cx="24" cy="17" r="4" fill="#555"/><path d="M32 35a8 8 0 1 0-16 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAddressBook0)"/>`),
		g.Group(children),
	)
}