package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushbullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640.5 768q-27 0-45.5-18.5T576.5 704V64q0-27 18.5-45.5T640.5 0q104 0 192.5 51.5t140 140t51.5 192.5T973 576.5t-140 140T640.5 768zm-192 0h-384q-26 0-45-19t-19-45V64q0-27 19-45.5T64.5 0h384q27 0 45.5 18.5T512.5 64v640q0 27-18.5 45.5T448.5 768z"/>`),
		g.Group(children),
	)
}