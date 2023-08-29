package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64zm-64 213.3H277.3V320h-42.7V213.3H128v-42.7h106.7V64h42.7v106.7H384v42.6z"/>`),
		g.Group(children),
	)
}