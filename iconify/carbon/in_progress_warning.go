package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InProgressWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.38 28h-6.762L24 21.236ZM24 18a1 1 0 0 0-.895.553l-5 10A1 1 0 0 0 19 30h10a1 1 0 0 0 .921-1.39l-5.026-10.057A1 1 0 0 0 24 18Z"/><path fill="currentColor" d="M18.746 22.8A9.999 9.999 0 1 1 14 4v10l6.097 6.097l1.22-2.44A2.985 2.985 0 0 1 24 16h1.82A11.993 11.993 0 1 0 14 26a11.934 11.934 0 0 0 3.394-.497Z"/>`),
		g.Group(children),
	)
}