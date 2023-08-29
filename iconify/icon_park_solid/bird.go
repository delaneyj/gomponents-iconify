package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBird0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m9 14l-5 6.07S5.85 27.035 11 32c9.89 9.533 24.334 3.303 30-1c5.357-4.37 2.717-5.332 1-5l-5 1c9.065-14.301 6.575-15.828 4-15l-9 4c-5.769 3.177-8.5 1.5-10 0l-3-3c-4.5-4-8.97-.16-10 1Z"/><circle cx="14" cy="20" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBird0)"/>`),
		g.Group(children),
	)
}