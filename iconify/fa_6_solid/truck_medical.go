package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 48C0 21.5 21.5 0 48 0h320c26.5 0 48 21.5 48 48v48h50.7c17 0 33.3 6.7 45.3 18.7l77.3 77.3c12 12 18.7 28.3 18.7 45.3V352c17.7 0 32 14.3 32 32s-14.3 32-32 32h-32c0 53-43 96-96 96s-96-43-96-96H256c0 53-43 96-96 96s-96-43-96-96H48c-26.5 0-48-21.5-48-48V48zm416 208h128v-18.7L466.7 160H416v96zM160 464a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm368-48a48 48 0 1 0-96 0a48 48 0 1 0 96 0zM176 80v48h-48c-8.8 0-16 7.2-16 16v32c0 8.8 7.2 16 16 16h48v48c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16v-48h48c8.8 0 16-7.2 16-16v-32c0-8.8-7.2-16-16-16h-48V80c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16z"/>`),
		g.Group(children),
	)
}