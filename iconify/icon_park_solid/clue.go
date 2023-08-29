package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClue0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M37 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM11 42a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="M37 16v19.504A6.496 6.496 0 0 1 30.504 42v0a6.496 6.496 0 0 1-6.496-6.496v-23A6.504 6.504 0 0 0 17.504 6v0A6.504 6.504 0 0 0 11 12.504V32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClue0)"/>`),
		g.Group(children),
	)
}