package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherSnowTwoCloudSnowSnowfallOvercastWeatherPrecipitationMeteorology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="currentColor" d="M3.25 10a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm-2 2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm6-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM5 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm5.75-2.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75ZM9 12.5a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Zm3.75 0a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 8a2.5 2.5 0 0 0 0-5a2.54 2.54 0 0 0-1.57.55A3.75 3.75 0 1 0 5.25 8Z"/>`),
		g.Group(children),
	)
}