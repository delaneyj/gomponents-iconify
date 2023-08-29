package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardboardBoxClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 73.32L74.6 155.1l82.3 37.1l173.2-85.5L256 73.32zm95.4 42.98l-173.2 85.5l77.8 35.1l181.4-81.8l-86-38.8zM61.7 169v182L247 434.6v-182L61.7 169zm388.6 0L265 252.6v182L450.3 351V169z"/>`),
		g.Group(children),
	)
}