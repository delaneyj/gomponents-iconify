package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.717 13.77A5.917 5.917 0 0 0 23 9c0-3.533-2.878-6-7-6a6.627 6.627 0 0 0-7 7v21h2v-5.88a8.126 8.126 0 0 0 6 2.48a7.383 7.383 0 0 0 7.653-7.6a6.636 6.636 0 0 0-3.936-6.23ZM17 25.6a5.678 5.678 0 0 1-6-5.6V10a4.686 4.686 0 0 1 5-5c3.037 0 5 1.57 5 4a3.873 3.873 0 0 1-4 4v2c3.434 0 5.653 1.963 5.653 5A5.39 5.39 0 0 1 17 25.6Z"/>`),
		g.Group(children),
	)
}