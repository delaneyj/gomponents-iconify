package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M656.5 907.5Q583 984 484 1011t-198 0t-172.5-103.5t-100-179t0-204.5T113 345L385 0l272 345q73 77 99.5 179t0 204.5t-100 179zM204 834q8 9 11.5 13.5T231 858t26 6q31 0 46.5-17.5T319 806q0-28-23-80.5T249.5 613T226 507q0-28 24-74.5t47-79.5l24-33h-32l-85 97q-75 93-75 208.5T204 834z"/>`),
		g.Group(children),
	)
}