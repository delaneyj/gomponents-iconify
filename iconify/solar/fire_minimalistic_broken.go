package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16 19.997A8.378 8.378 0 0 1 12 21c-4.418 0-8-3.356-8-7.496c0-1.131.15-2.192.405-3.176m14.771 6.493c.528-1 .824-2.126.824-3.317c0-3.741-2.035-6.666-3.438-8.06c-.26-.258-.694-.144-.84.189c-.748 1.69-2.304 4.123-4.293 4.123c-1.232.165-3.112-.888-1.594-6.107c.137-.47-.365-.848-.749-.534A14.639 14.639 0 0 0 6 6.647"/>`),
		g.Group(children),
	)
}