package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceWeatherLightningTwoCloudBoltStormWeatherThunderMeteorologyLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.78 9.38A2.5 2.5 0 1 0 9.5 5h0a4.5 4.5 0 1 0-5.75 4.32"/><path d="m8 8l-2 2.5h3l-2 3"/></g>`),
		g.Group(children),
	)
}