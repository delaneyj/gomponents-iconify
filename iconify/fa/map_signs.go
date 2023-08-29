package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSigns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1713 297q10 10 10 23t-10 23l-141 141q-28 28-68 28H160q-26 0-45-19t-19-45V192q0-26 19-45t45-19h576V64q0-26 19-45t45-19h128q26 0 45 19t19 45v64h512q40 0 68 28zm-977 919h256v512q0 26-19 45t-45 19H800q-26 0-45-19t-19-45v-512zm832-448q26 0 45 19t19 45v256q0 26-19 45t-45 19H224q-40 0-68-28L15 983Q5 973 5 960t10-23l141-141q28-28 68-28h512V576h256v192h576z"/>`),
		g.Group(children),
	)
}