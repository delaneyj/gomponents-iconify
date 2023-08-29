package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M283 128V64q-4-3-11.5-8.5T239 39t-50.5-15.5t-64 .5T48 49V21q0-9-6-15T27 0Q17 0 11 6T5 21v491h43V320q64-44 149-13v98q10 6 24 10.5t23.5 7T273 426t25 1h49q71 0 121-47l7-7V81l-37 32q-3 2-8 5t-21 9.5t-33 10t-42.5 2T283 128zM48 277V98q110-64 192-13v192q-98-50-192 0zm384 77q-118 69-192 30v-64h43V171q86 26 149-2v185z"/>`),
		g.Group(children),
	)
}