package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amoledpix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.003 6.214S6.665 38.312 6.141 40.267C5.618 42.22 7.176 42.5 7.774 42.5h8.711c.396 0 .732.012 1.55-1.495c.758-1.398 5.99-18.076 5.99-18.076l5.86 18.076s.459 1.495 1.425 1.495h8.895c.58 0 2.168-.372 1.662-2.233c-.506-1.858-12.984-34.053-12.984-34.053s-.2-.714-1.484-.714h-6.867c-1.27 0-1.53.714-1.53.714h.001Z"/>`),
		g.Group(children),
	)
}