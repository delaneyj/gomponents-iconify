package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagScout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M27 0Q17 0 11 6T5 21v491h43V384q73-2 136-18t105.5-40.5t76.5-54t54-58.5t33-52.5t17-38.5l5-15q-65 17-129 19t-110-5.5T147.5 99T84 74T48 53V21q0-9-6-15T27 0zm21 235v-34l21 6v49q-17-6-21-9v-12zm363-86q0 1-13 20t-16.5 24t-18.5 24.5t-24 27t-28.5 24t-36 24t-42.5 19t-52 16.5t-60.5 9t-71.5 4v-47h4q24 5 42-11.5t18-41.5v-26q0-17-10-31.5T76 164l-28-8v-54q88 39 140 49q99 19 223-2z"/>`),
		g.Group(children),
	)
}