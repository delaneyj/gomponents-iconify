package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClubsCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512zM43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21V64zm106 235q12 0 22-7v49h-22v22h86v-22h-22v-49q10 7 22 7q17 0 29.5-13t12.5-30t-12.5-30t-29.5-13q0-17-13-29.5T192 171t-30 12.5t-13 29.5q-17 0-29.5 13T107 256t12.5 30t29.5 13z"/>`),
		g.Group(children),
	)
}