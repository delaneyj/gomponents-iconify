package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scalpel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m486.594 19.438l-212.53 169.468c11.276 6.004 20.268 16.362 25.436 28.156a54.197 54.197 0 0 1 4 14.094c80.446 34.536 191.193-106.27 183.094-211.72zM252.03 202.125c-1.268.034-2.38.253-4.53.75c-4.932 1.14-8.54 2.213-12.03 5L18.22 382.063v27.78l233.218-184.937l11.625 14.656L18.22 433.656v32.53l255.81-204.155c13.41-10.704 14.012-24.534 8.345-37.467c-5.662-12.923-18.25-22.68-30.313-22.438h-.03z"/>`),
		g.Group(children),
	)
}