package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SatelliteRadar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 14h-2A10.011 10.011 0 0 0 18 4V2a12.014 12.014 0 0 1 12 12Z"/><path fill="currentColor" d="M26 14h-2a6.007 6.007 0 0 0-6-6V6a8.01 8.01 0 0 1 8 8zM16 28v-3.04a9.912 9.912 0 0 0 7.318-2.208a1.848 1.848 0 0 0 .678-1.334a1.8 1.8 0 0 0-.524-1.36L18.414 15L21 12.414L19.586 11L17 13.586l-5.058-5.059a1.815 1.815 0 0 0-1.36-.523a1.845 1.845 0 0 0-1.334.679a9.957 9.957 0 0 0-.513 11.95L6.28 28H2v2h28v-2zm-5.32-17.906L21.906 21.32A8.001 8.001 0 0 1 10.68 10.094zM14 28H8.387l1.876-5.627A9.99 9.99 0 0 0 14 24.543z"/>`),
		g.Group(children),
	)
}