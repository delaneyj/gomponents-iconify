package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 5a3.75 3.75 0 1 1 7.5 0v6.348c0 .072.04.18.158.274a6.25 6.25 0 1 1-7.816 0c.118-.095.158-.202.158-.274V5ZM12 2.75A2.25 2.25 0 0 0 9.75 5v6.348c0 .603-.31 1.116-.72 1.444a4.75 4.75 0 1 0 5.939 0c-.409-.328-.719-.84-.719-1.444V5A2.25 2.25 0 0 0 12 2.75Zm0 1.5a.75.75 0 0 1 .75.75v8.337a3.251 3.251 0 1 1-1.5 0V5a.75.75 0 0 1 .75-.75Zm0 10.5a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}