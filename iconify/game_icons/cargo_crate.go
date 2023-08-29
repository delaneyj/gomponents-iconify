package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CargoCrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247 16v89h9c14.9 0 23 11.7 23.8 23.6c.4 6-1.2 11.5-4.7 15.4c-3.5 4-9.1 7-19.1 7c-5 0-11.1-2.8-15.6-7.4c-4.6-4.5-7.4-10.6-7.4-15.6h-18c0 11 5.2 20.9 12.6 28.4c7.5 7.4 17.4 12.6 28.4 12.6c14 0 25.4-5 32.5-13c7.2-8.1 10-18.6 9.3-28.6c-1.2-17.5-13.4-35.18-32.8-39.42V16h-18zm-47.9 140.5L61.34 247h32.7l114.86-75.5l-9.8-15zm113.8 0l-9.8 15L418 247h32.6l-137.7-90.5zM41 265v222h430V265H41zm38 23h18v176H79V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288zm48 0h18v176h-18V288z"/>`),
		g.Group(children),
	)
}