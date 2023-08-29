package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Testcafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m317.837 313.07l-17.668 17.547h70.61l70.61 61.5H70.61l70.67-61.5h52.943l-17.667-17.546h-52.943L0 418.467h512L388.447 313.07h-70.61zm159.725-175.64l-44.14-43.896l-185.418 184.445l-70.612-70.305l-44.139 43.896l114.75 114.2l229.56-228.34z"/>`),
		g.Group(children),
	)
}