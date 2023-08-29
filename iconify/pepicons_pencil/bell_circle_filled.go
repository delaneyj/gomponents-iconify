package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilBellCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M11 20a.5.5 0 0 1 1 0a1 1 0 1 0 2 0a.5.5 0 0 1 1 0a2 2 0 1 1-4 0Z"/><path fill-rule="evenodd" d="M20.5 17.5a2.96 2.96 0 0 0-1.5-2.575V12a5.5 5.5 0 0 0-5.5-5.5h-1A5.5 5.5 0 0 0 7 12v2.925A2.96 2.96 0 0 0 5.5 17.5a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2ZM18 15.558l.295.133l.055.024A1.96 1.96 0 0 1 19.5 17.5a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1c0-.769.45-1.467 1.15-1.784l.055-.025l.295-.133V12a4.5 4.5 0 0 1 4.5-4.5h1A4.5 4.5 0 0 1 18 12v3.558Z" clip-rule="evenodd"/><path d="M12.5 4.5a.5.5 0 0 1 1 0V7a.5.5 0 0 1-1 0V4.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilBellCircleFilled0)"/></g>`),
		g.Group(children),
	)
}