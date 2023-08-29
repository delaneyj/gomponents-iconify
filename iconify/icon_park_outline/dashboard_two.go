package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20a19.937 19.937 0 0 0-5.845-14.13A19.938 19.938 0 0 0 24 4A19.938 19.938 0 0 0 9.845 9.87A19.937 19.937 0 0 0 4 24c0 11.046 8.954 20 20 20ZM4 24h4M9.845 9.87l3.13 3.13M24 4v4m20 16h-4M38.155 9.87L35.025 13M24 20v12"/><path d="M39.852 36.196C36.195 40.942 30.455 44 24 44s-12.195-3.058-15.852-7.804A31.853 31.853 0 0 1 24 32a31.854 31.854 0 0 1 15.852 4.196Z"/></g>`),
		g.Group(children),
	)
}