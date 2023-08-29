package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpadesCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512zM43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21V64zm128 207v49h-22v21h86v-21h-22v-49q9 6 32.5 3t31.5-18q11-20-10.5-47T218 166l-26-17q-2 1-6 3.5t-15 10t-20 14.5t-21 18t-17 20t-9 21t3 20q8 15 31.5 18t32.5-3z"/>`),
		g.Group(children),
	)
}