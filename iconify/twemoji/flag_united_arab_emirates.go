package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagUnitedArabEmirates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#068241" d="M32 5H9v9h27V9a4 4 0 0 0-4-4z"/><path fill="#EEE" d="M9 14h27v8H9z"/><path fill="#141414" d="M9 31h23a4 4 0 0 0 4-4v-5H9v9z"/><path fill="#EC2028" d="M4 5a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h5V5H4z"/>`),
		g.Group(children),
	)
}