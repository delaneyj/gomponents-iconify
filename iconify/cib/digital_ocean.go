package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalOcean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.438 30.113h4.606v-4.6H5.438zm-3.851-8.45v3.85h3.85v-3.85zm29.138-10.607c-1.438-4.637-5.15-8.331-9.788-9.787C10.306-2.05.499 5.856.499 15.994h5.988c0-6.362 6.312-11.281 13.006-8.856a8.968 8.968 0 0 1 5.363 5.356c2.444 6.688-2.481 12.988-8.837 13v.019H16v5.988c10.163 0 18.05-9.8 14.725-20.444zM16.019 25.494v-5.956h-5.975v5.975H16v-.019z"/>`),
		g.Group(children),
	)
}