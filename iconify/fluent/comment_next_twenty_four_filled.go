package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentNextTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12.023 2.999A6.5 6.5 0 0 0 22 11.189L22 14.75A3.25 3.25 0 0 1 18.75 18h-5.785l-5.387 3.817A1 1 0 0 1 6 21.002V18h-.75A3.25 3.25 0 0 1 2 14.75v-8.5A3.25 3.25 0 0 1 5.25 3l6.773-.001zM17.5 1a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm.216 2.589l-.07.057l-.057.07a.5.5 0 0 0 0 .568l.057.07L19.292 6H14l-.09.008a.5.5 0 0 0-.402.402l-.008.09l.008.09a.5.5 0 0 0 .402.402L14 7h5.292l-1.646 1.646l-.057.07a.5.5 0 0 0 .695.695l.07-.057l2.541-2.548l.033-.048l.034-.067l.021-.063l.015-.082L21 6.5l-.003-.053l-.014-.075l-.03-.083l-.042-.074l-.045-.056l-2.512-2.513l-.07-.057a.5.5 0 0 0-.568 0z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}