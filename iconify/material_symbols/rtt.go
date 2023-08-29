package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rtt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.125 21l.4-2.55h2.325l2.05-12.9h-2.5l-.725 4.5h-2.65L9.15 3h12.975L21 10.05h-2.65l.7-4.5h-2.5l-2.05 12.9h2.325l-.4 2.55h-7.3ZM3.75 7l.325-2h3.75L7.5 7H3.75Zm-.625 4l.325-2H7.2l-.325 2h-3.75Zm-1.25 8l.3-2h6.25l-.3 2h-6.25Zm.625-4l.325-2h6.25l-.325 2H2.5Z"/>`),
		g.Group(children),
	)
}