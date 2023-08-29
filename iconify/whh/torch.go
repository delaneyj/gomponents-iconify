package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m506 548l-38 64q-7 11-20 19.5t-25 8.5h-39l-64 352q0 13-9 22.5t-23 9.5h-64q-13 0-22.5-9.5T192 992l-64-352H89q-11 0-24-8.5T44 612L7 548q-9-15-5.5-25.5T19 512h475q13 0 16.5 10.5T506 548zM399 448h-79q2-4 6.5-17.5T334 406t3-22q0-38-57-135q-8-9-17-19q-4 21-11 44t-28 58t-47 52q-25 16-13 64H99q-18-3-26.5-23.5T64 384q0-21 16-43.5t32-44.5t16-40q0-27-6.5-54T106 157t-18.5-32.5T71 103l-6-7q58 24 120 68q14-45 8.5-81.5T174 23L161 0q6 3 17 8t41.5 21.5t58 35.5t60 48t53.5 60q29 41 44 90.5t13 88.5q-1 40-49 96z"/>`),
		g.Group(children),
	)
}