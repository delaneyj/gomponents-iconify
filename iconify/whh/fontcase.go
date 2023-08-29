package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M979 1023h-31q-26 0-50.5-18.5T864 959l-21-64H566l-21 64q-8 27-32.5 45.5T462 1023h-31q-27 0-39.5-18.5T388 959l188-575q8-27 32.5-45.5T659 320h91q26 0 50.5 18.5T834 384l188 575q8 27-4.5 45.5T979 1023zM705 471l-97 296h194zM353 172L246 576h204l-42 127H212l-50 192q-9 27-33.5 45.5T78 959H47q-26 0-38.5-18.5T4 895L224 64q9-26 33.5-45T308 0h91q26 0 50.5 19T482 64l61 229l-87 266z"/>`),
		g.Group(children),
	)
}