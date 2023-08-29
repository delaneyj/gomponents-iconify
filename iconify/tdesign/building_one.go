package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 1.5v1.625c.396.106.752.277 1.065.51c.549.407.897.951 1.113 1.494c.247.615.344 1.277.364 1.871H11.5v2H10v1h12v12H3V9H1.5V7h1.458c.02-.594.117-1.256.363-1.871c.217-.543.565-1.087 1.114-1.494a3.118 3.118 0 0 1 1.065-.51V1.5h2ZM4.96 7h3.08c-.018-.389-.083-.789-.219-1.129c-.116-.29-.268-.497-.449-.63C7.201 5.112 6.937 5 6.5 5s-.7.113-.873.24c-.18.134-.332.34-.449.631c-.135.34-.2.74-.219 1.129ZM8 9H5v11h3V9Zm2 11h2v-5h6v5h2v-8H10v8Zm6 0v-3h-2v3h2Z"/>`),
		g.Group(children),
	)
}