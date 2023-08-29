package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octoloadertwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 583V441l256-106v354zm192-153l-128 53v58l128 53V430zM643 280L749 25l251 250l-256 106zm-202-24L335 0h354L583 256H441zM280 381L24 275L275 24l106 256zm-27-244L137 253l128 53l41-41zm3 446L0 689V335l256 106v142zM64 430v164l128-53v-58zm317 313L275 999L24 749l257-106zM253 887l53-128l-41-41l-128 53zm330-119l106 256H335l106-256h142zm11 192l-53-128h-58l-53 128h164zm150-317l256 106l-251 250l-106-256zm27 244l116-116l-128-53l-41 41z"/>`),
		g.Group(children),
	)
}