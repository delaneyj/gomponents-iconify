package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octoloaderfive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 583V441l256-107v355zM643 280L749 24l251 251l-256 106zm-202-24L335 0h354L583 256H441zM280 381L24 275L275 24l106 256zm-27-244L137 253l128 53l41-41zm3 446L0 689V334l256 107v142zM64 430v164l128-53v-58zm317 313L275 999L24 749l256-107zM253 887l53-128l-41-41l-128 53zm330-119l106 256H335l106-256h142zm160-126l257 107l-251 250l-106-256z"/>`),
		g.Group(children),
	)
}