package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherFogTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.745 19.5h8.501a.75.75 0 0 1 .102 1.493l-.102.007H7.745a.75.75 0 0 1-.102-1.493l.102-.007h8.501h-8.501ZM4.75 16.52h14.5a.75.75 0 0 1 .102 1.492l-.102.007H4.75a.75.75 0 0 1-.102-1.493l.102-.007ZM12 3.004c3.168 0 4.966 2.097 5.227 4.63h.08A3.687 3.687 0 0 1 21 11.318A3.687 3.687 0 0 1 17.307 15H6.693A3.687 3.687 0 0 1 3 11.318a3.687 3.687 0 0 1 3.693-3.683h.08c.262-2.55 2.058-4.63 5.227-4.63Z"/>`),
		g.Group(children),
	)
}