package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.462 6.81l3.284 13.616c.178.737 1.211.775 1.443.054l3.257-10.123l.586 2.096a.75.75 0 0 0 .722.548h3.494a.75.75 0 0 0 0-1.5h-2.925l-1.105-3.95c-.2-.717-1.208-.736-1.436-.028l-3.204 9.957L9.226 3.574c-.182-.757-1.255-.769-1.454-.016L5.67 11.501H2.75a.75.75 0 0 0 0 1.5h3.496a.75.75 0 0 0 .725-.558L8.462 6.81Z"/>`),
		g.Group(children),
	)
}