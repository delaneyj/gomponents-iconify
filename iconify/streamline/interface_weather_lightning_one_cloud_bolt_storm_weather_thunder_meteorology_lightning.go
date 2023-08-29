package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherLightningOneCloudBoltStormWeatherThunderMeteorologyLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.5 7L6 10.5h3l-2 3"/><path d="M11 8a2.5 2.5 0 1 0-.62-4.92a3.5 3.5 0 0 0-6.76 0A2.5 2.5 0 1 0 3 8h1.5"/></g>`),
		g.Group(children),
	)
}