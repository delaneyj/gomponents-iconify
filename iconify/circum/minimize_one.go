package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimizeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17.78a.5.5 0 0 1-1 0v-3.07l-6.08 6.08a.5.5 0 0 1-.71-.71c.29-.29.58-.57.87-.86C5.82 17.48 7.55 15.74 9.3 14H6.22a.5.5 0 0 1 0-1h4.28a.429.429 0 0 1 .34.14c.01.01.02.01.02.02a.384.384 0 0 1 .13.26ZM14.7 10h3.08a.5.5 0 0 1 0 1H13.5a.429.429 0 0 1-.34-.14c-.01-.01-.02-.01-.02-.02a.384.384 0 0 1-.13-.26L13 6.22a.5.5 0 0 1 1 0v3.07l.01-.01l6.07-6.07a.5.5 0 0 1 .71.71Z"/>`),
		g.Group(children),
	)
}