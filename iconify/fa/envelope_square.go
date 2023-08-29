package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960zm32 1056V620q-31 35-64 55q-34 22-132.5 85T932 859q-98 69-164 69t-164-69q-47-32-142-92.5T320 674q-12-8-33-27t-31-27v436q0 40 28 68t68 28h832q40 0 68-28t28-68zm0-573q0-41-27.5-70t-68.5-29H352q-40 0-68 28t-28 68q0 37 30.5 76.5T354 621q47 32 137.5 89T621 793q3 2 17 11.5t21 14t21 13t23.5 13T725 854t22.5 7.5T768 864t20.5-2.5T811 854t21.5-9.5t23.5-13t21-13t21-14t17-11.5l267-174q35-23 66.5-62.5T1280 483z"/>`),
		g.Group(children),
	)
}