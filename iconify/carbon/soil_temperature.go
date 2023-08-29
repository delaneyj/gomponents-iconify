package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilTemperature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="5" cy="13" r="1" fill="currentColor"/><circle cx="11" cy="19" r="1" fill="currentColor"/><circle cx="15" cy="25" r="1" fill="currentColor"/><circle cx="17" cy="15" r="1" fill="currentColor"/><circle cx="13" cy="11" r="1" fill="currentColor"/><circle cx="9" cy="27" r="1" fill="currentColor"/><circle cx="3" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M25 30a4.986 4.986 0 0 1-3-8.98V15a3 3 0 0 1 6 0v6.02A4.986 4.986 0 0 1 25 30zm0-16a1.001 1.001 0 0 0-1 1v7.13l-.497.29A2.968 2.968 0 0 0 22 25a3 3 0 0 0 6 0a2.968 2.968 0 0 0-1.503-2.581L26 22.13V15a1.001 1.001 0 0 0-1-1zM2 6h28v2H2z"/>`),
		g.Group(children),
	)
}