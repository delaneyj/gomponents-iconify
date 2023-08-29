package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 208a4 4 0 0 1-4 4H32a4 4 0 0 1 0-8h192a4 4 0 0 1 4 4ZM29.2 157.12a12 12 0 0 1 1.51-12.83l81.66-100.83a20.1 20.1 0 0 1 31.26 0l81.66 100.83a12.1 12.1 0 0 1-9.37 19.71H40.08a12 12 0 0 1-10.88-6.88Zm7.22-3.44a4 4 0 0 0 3.66 2.32h175.84a4 4 0 0 0 3.66-2.32a4 4 0 0 0-.51-4.36L137.41 48.5a12.09 12.09 0 0 0-18.82 0L36.93 149.32a4 4 0 0 0-.51 4.36Z"/>`),
		g.Group(children),
	)
}