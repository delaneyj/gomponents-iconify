package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilMoistureGlobal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 30a4.735 4.735 0 0 1-4.207-7.35l3.368-5.194a1 1 0 0 1 1.678 0l3.315 5.11A4.975 4.975 0 0 1 30 25.223A4.898 4.898 0 0 1 25 30zm0-10.163l-2.508 3.866A2.926 2.926 0 1 0 28 25.223a2.96 2.96 0 0 0-.544-1.603zM11 24H9A17.02 17.02 0 0 1 26 7h4v2h-4a15.017 15.017 0 0 0-15 15z"/><path fill="currentColor" d="M16 24h-2a12.014 12.014 0 0 1 12-12h4v2h-4a10.011 10.011 0 0 0-10 10Z"/><path fill="currentColor" d="M26 2a21.863 21.863 0 0 0-12 3.576A21.856 21.856 0 0 0 2 2v2a19.873 19.873 0 0 1 10.262 2.835a22.151 22.151 0 0 0-2.315 2.144A17.001 17.001 0 0 0 2 7v2a15.014 15.014 0 0 1 6.62 1.537a22.077 22.077 0 0 0-1.702 2.537A12.032 12.032 0 0 0 2 12v2a9.884 9.884 0 0 1 4 .85A21.861 21.861 0 0 0 4 24h2A20.023 20.023 0 0 1 26 4h4V2Z"/>`),
		g.Group(children),
	)
}