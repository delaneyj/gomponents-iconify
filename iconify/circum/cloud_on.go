package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.917 13.484a4.381 4.381 0 0 0-5.19-4.26a6.281 6.281 0 0 0-11.75 2.19a3.237 3.237 0 0 0-2.66 2a3.433 3.433 0 0 0 .82 3.74c1.12 1.03 2.54.89 3.94.89h10.15a4.514 4.514 0 0 0 4.69-4.32Zm-4.65 3.56c-1.19.01-2.38 0-3.56 0c-2.75 0-5.49.06-8.23 0a2.383 2.383 0 0 1-2.33-1.73a2.333 2.333 0 0 1 2.28-2.94a.515.515 0 0 0 .5-.5a5.3 5.3 0 0 1 10.11-1.81a.5.5 0 0 0 .56.23a3.366 3.366 0 0 1 4.33 3.32a3.489 3.489 0 0 1-3.66 3.43Z"/>`),
		g.Group(children),
	)
}