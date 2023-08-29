package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castlevania(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.624 5.842L24.062 3.5l-1.56 2.342s1.56-.063 1.56 1.811c0-1.874 1.562-1.811 1.562-1.811Zm4.527 11.522v-5.246L28.56 9.59s-4.497.828-4.497-1.936c0 2.764-4.496 1.936-4.496 1.936l-1.717 2.53v22.16l1.686 2.342s2.56-.5 3.414 1.384a2.154 2.154 0 0 1-.479 2.363l1.592 4.132l1.593-4.132a2.154 2.154 0 0 1-.479-2.363c.854-1.884 3.414-1.384 3.414-1.384l1.561-2.342v-5.746l-3.622 1.874v2.872c0 .948-1.104 1.717-2.467 1.717s-2.591-.769-2.591-1.717V13.12c0-.949 1.23-1.718 2.591-1.718s2.467.77 2.467 1.718v2.185Z"/>`),
		g.Group(children),
	)
}