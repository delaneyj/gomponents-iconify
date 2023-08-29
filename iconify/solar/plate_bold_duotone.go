package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlateBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M1 12c0-3.75 0-5.625.955-6.939A5 5 0 0 1 3.06 3.955C4.375 3 6.251 3 10 3h4c3.75 0 5.625 0 6.939.955a5 5 0 0 1 1.106 1.106C23 6.375 23 8.251 23 12c0 3.75 0 5.625-.955 6.939a5 5 0 0 1-1.106 1.106C19.625 21 17.749 21 14 21h-4c-3.75 0-5.625 0-6.939-.955a5 5 0 0 1-1.106-1.106C1 17.625 1 15.749 1 12Z" opacity=".5"/><path d="M12.75 3a.75.75 0 0 0-1.5 0v2a.75.75 0 0 0 1.5 0V3ZM8 9.75a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5H8Zm0 3.5a.75.75 0 0 0 0 1.5h5.5a.75.75 0 0 0 0-1.5H8Z"/></g>`),
		g.Group(children),
	)
}