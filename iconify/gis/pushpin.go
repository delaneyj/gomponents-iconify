package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M37.234 0a5.87 5.87 0 0 0-5.87 5.87a5.87 5.87 0 0 0 5.57 5.86l-1.452 28.32a11.74 11.74 0 0 0-9.987 11.606h19.28l3.829 43.1a1.612 10.476 0 0 0 2.792 0l3.83-43.1h19.279a11.74 11.74 0 0 0-9.987-11.606l-1.451-28.32a5.87 5.87 0 0 0 5.569-5.86A5.87 5.87 0 0 0 62.766 0Z" color="currentColor"/>`),
		g.Group(children),
	)
}