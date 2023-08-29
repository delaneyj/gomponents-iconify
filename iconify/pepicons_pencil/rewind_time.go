package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.604 5.45a6.44 6.44 0 0 0-1.883 5.278a.5.5 0 0 1-.994.105a7.44 7.44 0 0 1 2.175-6.096c2.937-2.897 7.675-2.85 10.582.098c2.907 2.947 2.888 7.685-.05 10.582a7.425 7.425 0 0 1-5.097 2.142a7.527 7.527 0 0 1-2.14-.271a.5.5 0 0 1 .266-.964a6.524 6.524 0 0 0 1.856.235a6.424 6.424 0 0 0 4.413-1.854c2.541-2.506 2.562-6.61.04-9.168c-2.522-2.558-6.627-2.594-9.168-.088Z"/><path d="M3.594 11.363a.5.5 0 0 1-.706.04l-1.72-1.53a.5.5 0 1 1 .664-.746l1.72 1.53a.5.5 0 0 1 .042.706Z"/><path d="M2.82 11.3a.5.5 0 0 0 .7.1l2-1.5a.5.5 0 1 0-.6-.8l-2 1.5a.5.5 0 0 0-.1.7ZM10 6.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V7a.5.5 0 0 1 .5-.5Z"/><path d="M13.5 10.5a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z"/></g>`),
		g.Group(children),
	)
}