package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSymbolNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHospitalSymbolNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM3.97 24.291c0-11.046 8.954-20 20-20c11.045 0 20 8.954 20 20s-8.955 20-20 20c-11.046 0-20-8.954-20-20Zm16.021 1.757v7h-4v-18h4v7h8v-7h4v18h-4v-7h-8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHospitalSymbolNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}