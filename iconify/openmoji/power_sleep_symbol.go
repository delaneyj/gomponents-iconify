package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSleepSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M53.537 45.624c-7.807 5.525-18.688 4.792-25.68-2.2c-6.434-6.435-7.567-16.165-3.398-23.76C19.34 23.286 16 29.253 16 36c0 11.046 8.954 20 20 20c7.557 0 14.135-4.191 17.537-10.376Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}