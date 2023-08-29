package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneVideoCall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPhoneVideoCall0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M14.376 9.794a2 2 0 0 1 1.748 1.03l2.447 4.406a2 2 0 0 1 .04 1.865l-2.357 4.714s.683 3.511 3.541 6.37c2.859 2.858 6.358 3.53 6.358 3.53l4.713-2.357a2 2 0 0 1 1.867.04l4.42 2.458a2 2 0 0 1 1.027 1.748v5.073c0 2.584-2.4 4.45-4.848 3.624c-5.028-1.696-12.832-4.927-17.78-9.873c-4.946-4.947-8.176-12.752-9.873-17.78c-.826-2.448 1.04-4.848 3.624-4.848h5.073Z"/><path stroke-linecap="round" d="M39 15H27V5h12v3l5-2v8l-5-2v3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPhoneVideoCall0)"/>`),
		g.Group(children),
	)
}