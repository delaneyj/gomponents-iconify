package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.012 2.004C6.492 2 2.013 6.474 2.01 11.994S6.48 21.993 12 21.996a9.996 9.996 0 0 0 10.002-9.99c.003-5.52-4.47-9.999-9.99-10.002zm0 18.992A8.996 8.996 0 1 1 12 3.004a8.996 8.996 0 0 1 .012 17.992zM12.707 12l3.182-3.182a.5.5 0 0 0-.707-.707L12 11.293L8.818 8.111a.5.5 0 0 0-.707.707L11.293 12l-3.182 3.182a.5.5 0 0 0 .707.707L12 12.707l3.182 3.182a.5.5 0 0 0 .707-.707L12.707 12z"/>`),
		g.Group(children),
	)
}