package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurntableLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M2 12c0-4.714 0-7.071 1.172-8.536C4.343 2 6.229 2 10 2h4c3.771 0 5.657 0 6.828 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.172 8.535C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.465C2 19.072 2 16.714 2 12Z"/><rect width="8" height="8" x="6" y="8" stroke="currentColor" stroke-width="1.5" rx="4"/><path fill="currentColor" d="M14.665 16.33a.75.75 0 1 0 .67 1.34l-.67-1.34Zm1.23.223l-.336-.671l.335.67Zm-.56 1.118l.895-.447l-.671-1.342l-.894.447l.67 1.342Zm2.415-2.907V10.5h-1.5v4.264h1.5Zm-1.52 2.46a2.75 2.75 0 0 0 1.52-2.46h-1.5c0 .473-.268.906-.691 1.118l.67 1.342ZM17 9.75a.75.75 0 0 1-.75-.75h-1.5A2.25 2.25 0 0 0 17 11.25v-1.5Zm.75-.75a.75.75 0 0 1-.75.75v1.5A2.25 2.25 0 0 0 19.25 9h-1.5ZM17 8.25a.75.75 0 0 1 .75.75h1.5A2.25 2.25 0 0 0 17 6.75v1.5Zm0-1.5A2.25 2.25 0 0 0 14.75 9h1.5a.75.75 0 0 1 .75-.75v-1.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}