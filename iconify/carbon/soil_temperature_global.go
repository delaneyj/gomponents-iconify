package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilTemperatureGlobal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 30a4.986 4.986 0 0 1-3-8.98V15a3 3 0 0 1 6 0v6.02A4.986 4.986 0 0 1 25 30zm0-16a1.001 1.001 0 0 0-1 1v7.13l-.497.29A2.968 2.968 0 0 0 22 25a3 3 0 0 0 6 0a2.968 2.968 0 0 0-1.503-2.581L26 22.13V15a1.001 1.001 0 0 0-1-1zm-9 10h-2a11.89 11.89 0 0 1 3.764-8.727l1.373 1.454A9.909 9.909 0 0 0 16 24z"/><path fill="currentColor" d="M11 24H9A17.02 17.02 0 0 1 26 7h4v2h-4a15.017 15.017 0 0 0-15 15Z"/><path fill="currentColor" d="M26 2a21.863 21.863 0 0 0-12 3.576A21.856 21.856 0 0 0 2 2v2a19.873 19.873 0 0 1 10.262 2.835a22.151 22.151 0 0 0-2.315 2.144A17.001 17.001 0 0 0 2 7v2a15.014 15.014 0 0 1 6.62 1.537a22.077 22.077 0 0 0-1.702 2.537A12.032 12.032 0 0 0 2 12v2a9.884 9.884 0 0 1 4 .85A21.861 21.861 0 0 0 4 24h2A20.023 20.023 0 0 1 26 4h4V2Z"/>`),
		g.Group(children),
	)
}