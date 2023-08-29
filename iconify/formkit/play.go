package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.5 13.62c-.24 0-.47-.06-.69-.17c-.5-.26-.81-.77-.81-1.33V3.98c0-.56.31-1.07.81-1.33s1.1-.22 1.56.11l5.77 4.07c.4.28.63.74.63 1.23s-.24.94-.63 1.23l-5.77 4.06c-.26.18-.56.28-.86.28Zm0-10.14c-.08 0-.16.02-.23.06c-.17.09-.27.25-.27.44v8.14c0 .19.1.36.27.44c.17.09.36.07.52-.04l5.77-4.07a.507.507 0 0 0 0-.82L2.79 3.57a.521.521 0 0 0-.29-.09Z"/>`),
		g.Group(children),
	)
}