package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satisfyer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="19.733" cy="14.231" r="3.646" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.165" cy="26.259" r="3.646" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.56 12.95l11.985 11.99c1.332 1.332.122 12.956-6.995 12.956c-8.32-.001-10.05-2.673-10.05-2.673l5.237-5.24c2.873 2.874 6.148-.84 4.655-2.333c-.96-.807-6.514-6.516-6.514-6.516Zm31.794-2.845h-7.963a5.739 5.739 0 0 0-5.033 4.545l-6.1 22.773h7.713l4.489-16.756s.65-3.16 3.158-3.16H43.5Z"/>`),
		g.Group(children),
	)
}