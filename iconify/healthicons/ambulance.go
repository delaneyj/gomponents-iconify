package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M32 12.5a1 1 0 1 0 2 0V11a1 1 0 1 0-2 0v1.5Z"/><path fill-rule="evenodd" d="M4 16a2 2 0 0 1 2-2h22a2 2 0 0 1 1.732 1H30v4h1v-1h1v-1a1 1 0 1 1 2 0v1h1v1h.718a3 3 0 0 1 2.035.796l5.282 4.875A3 3 0 0 1 44 26.876V35h-5.126a4.002 4.002 0 0 1-7.748 0H15.874a4.002 4.002 0 0 1-7.748 0H4V16Zm38 12H30v5h1.126a4.002 4.002 0 0 1 7.748 0H42v-5Zm-12-2h11.526l-5.13-4.735a1 1 0 0 0-.678-.265H30v5ZM12 36a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm25-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM13 22h3v-3h2v3h3v2h-3v3h-2v-3h-3v-2Z" clip-rule="evenodd"/><path d="M36.5 17a1 1 0 0 1 1-1H39a1 1 0 1 1 0 2h-1.5a1 1 0 0 1-1-1Zm-.672-4.241a1 1 0 1 0 1.343 1.482l.915-.828a1 1 0 0 0-1.343-1.482l-.915.828Z"/></g>`),
		g.Group(children),
	)
}