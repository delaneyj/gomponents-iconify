package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24h0a3.285 3.285 0 0 1-3.285 3.285H29.69L24 17.428l5.262-9.114a3.285 3.285 0 0 1 4.488-1.202h0a3.285 3.285 0 0 1 1.203 4.488l-5.262 9.115h10.524A3.285 3.285 0 0 1 43.5 24h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.75 40.888h0a3.285 3.285 0 0 1-4.488-1.202L24 30.572l-5.262 9.114a3.285 3.285 0 0 1-4.488 1.202h0a3.285 3.285 0 0 1-1.203-4.488l5.261-9.115H29.69l5.262 9.115a3.285 3.285 0 0 1-1.202 4.488h0ZM24 17.428l-1.92 3.325l-3.771-.038l-5.262-9.115a3.285 3.285 0 0 1 1.203-4.488h0a3.285 3.285 0 0 1 4.488 1.202L24 17.428Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.08 20.753l-3.771 6.532H7.785A3.285 3.285 0 0 1 4.5 24h0a3.285 3.285 0 0 1 3.285-3.285l14.295.038Zm-3.771 6.532L24 17.428l5.691 9.857H18.309z"/>`),
		g.Group(children),
	)
}