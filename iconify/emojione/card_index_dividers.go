package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIndexDividers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#89664c" d="M63.4 24.7H47.7V59l16-32.4c.5-1 .4-1.8-.3-1.9"/><path fill="#aa7d5e" d="M60.8 21.3H45.1L47.7 59l13.4-35.8c.5-1 .3-1.8-.3-1.9"/><path fill="#d3976e" d="M56.1 17.9h-8.4V59l9.4-38.8c.3-1.3-.2-2.3-1-2.3"/><path fill="#f2bc97" d="M52.2 13h-8.4l3.9 46l5.4-43.7c.3-1.3-.1-2.3-.9-2.3"/><path fill="#ffce31" d="M45.8 11h-22V7.6C23.8 6.2 23 5 22 5H1.9C.8 5 0 6.2 0 7.6V59h47.7V13.6c0-1.5-.9-2.6-1.9-2.6z"/><g fill="#fff"><circle cx="10.8" cy="54.7" r="2.5"/><path d="M23.9 52.2c-1.4 0-2.5 1.1-2.5 2.5s1.1 2.5 2.5 2.5s2.5-1.1 2.5-2.5s-1.1-2.5-2.5-2.5"/><circle cx="36.9" cy="54.7" r="2.5"/></g>`),
		g.Group(children),
	)
}