package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPassportOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M13 40H9a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v4"/><rect width="34" height="28" x="13" y="44" fill="#555" rx="2" transform="rotate(-90 13 44)"/><circle cx="27" cy="27" r="8"/><path stroke-linecap="round" d="M35 28s-3-.5-5 1c-2.001 1.5-2.599 5.102-2 6m-4-8c0 2-2 4-4 4m11-11s0 4-3 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPassportOne0)"/>`),
		g.Group(children),
	)
}