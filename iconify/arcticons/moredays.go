package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moredays(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.488 23.662s-1.15-7.584 2.456-10.469s7.214.175 7.833 2.526a46.155 46.155 0 0 1 .776 7.308s1.254-6.316 3.453-8.528s4.003-2.888 6.561-1.066c2.203 1.57 2.138 2.385 1.938 10.792m.126 1.985a6.669 6.669 0 0 1-3.605 7.405c-4.831 2.489-8.955 3.082-12.886 1.447s-5.732-2.887-6.516-6.088s-.216-2.78.145-2.606a54.83 54.83 0 0 0 11.093 1.17c5.964.11 11.896.657 11.77-1.328Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}