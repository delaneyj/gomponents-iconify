package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoatHorizon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M313 80v48h18V80zm-78.5 29.2l-17 5.6l16 48l17-5.6zm175 0l-16 48l17 5.6l16-48zM322 167c-71.9 0-130.9 55.5-136.6 126h127.1c5.7-18 6.2-37.7 8.4-54.8c10.5 6.9 21.1 22.4 26 37.7c4.9-22.7 5.7-46.6 8.2-67.6c20.3 14.8 40.9 56.5 37 84.7h66.5c-5.7-70.5-64.7-126-136.6-126zM20 311v18h472v-18zm190 38v18h224v-18zm16 38v18h192v-18zm32 38v18h128v-18zm42 38v18h44v-18z"/>`),
		g.Group(children),
	)
}