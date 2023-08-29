package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 896q-126 0-232-65T181 658l50 31q13 10 34 10.5t38-11t17-31.5l-1-418q0-21-17-32t-38-10.5t-33 10.5l-50 30q58-108 163.5-172.5T576 0q91 0 174 35.5T893 131t95.5 143t35.5 174t-35.5 174T893 765t-143 95.5T576 896zM256 297v302q-13 9-30.5 9t-30.5-9L13 469q-13-9-13-21.5T13 426l182-129q12-9 30-9t31 9z"/>`),
		g.Group(children),
	)
}