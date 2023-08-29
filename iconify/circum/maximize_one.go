package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.065 16.16a.5.5 0 0 1 1 0v3.07l.01-.01l6.07-6.07a.5.5 0 0 1 .71.71c-.29.29-.58.57-.87.86c-1.74 1.74-3.47 3.48-5.21 5.22h3.07a.5.5 0 0 1 0 1h-4.28a.429.429 0 0 1-.34-.14c-.01-.01-.02-.01-.02-.02a.384.384 0 0 1-.13-.26c-.009-.078-.01-4.36-.01-4.36Zm17.87-12.6v4.28a.5.5 0 0 1-1 0V4.77l-.01.01q-3.045 3.03-6.07 6.07a.5.5 0 0 1-.71-.71c.29-.29.58-.57.86-.86c1.75-1.74 3.48-3.48 5.22-5.22h-3.07a.5.5 0 0 1 0-1h4.28a.429.429 0 0 1 .34.14c.01.01.02.01.02.02a.429.429 0 0 1 .14.34Z"/>`),
		g.Group(children),
	)
}