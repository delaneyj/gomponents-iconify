package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wordpress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.06 1C5.3 1 2.71 2.62 1.52 5.2h3.35v.37c-1.03-.12-.94.41-.79.97l2.2 5.48l1.61-3.77l-.74-1.77c-.27-.47-.27-1-1.38-.95l.02-.34h4.91v.39c-.57.03-1.27-.14-.9 1.03l1.99 5.19l1.09-3.03c.36-1.27.02-2.49-.4-3.69c-.23-.8-.13-1.47.93-1.74C11.8 1.71 9.91.99 8.08.99Zm6.68 4.79l-3.18 8.26c.88-.76 1.78-1.03 2.88-3.13c.73-1.74.66-3.36.3-5.13Zm-13.55.13c-.31 2.44-.28 3.15.49 5.14c1 1.91 1.78 2.31 2.88 3.06L1.19 5.93Zm7 3.09l-2.36 5.66c1.66.47 3.14.41 4.49-.11L8.19 9.01Z"/>`),
		g.Group(children),
	)
}