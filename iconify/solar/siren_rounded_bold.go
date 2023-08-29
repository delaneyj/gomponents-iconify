package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SirenRoundedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.75 2a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0V2Z"/><path fill-rule="evenodd" d="M2 21.25h2V16a8 8 0 1 1 16 0v5.25h2a.75.75 0 0 1 0 1.5H2a.75.75 0 0 1 0-1.5Zm10.75-2.45a1.5 1.5 0 1 0-1.5 0v2.45h1.5V18.8Zm.845-7.581a.75.75 0 0 1 .977-.414a4.764 4.764 0 0 1 2.623 2.623a.75.75 0 0 1-1.39.563a3.264 3.264 0 0 0-1.796-1.796a.75.75 0 0 1-.414-.976Z" clip-rule="evenodd"/><path d="M21.53 5.47a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l1.5-1.5a.75.75 0 0 1 1.06 0Zm-18 0a.75.75 0 0 0-1.06 1.06l1.5 1.5a.75.75 0 0 0 1.06-1.06l-1.5-1.5Z"/></g>`),
		g.Group(children),
	)
}