package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEgg0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="10" fill="#555" stroke-linecap="round" stroke-linejoin="round"/><path d="M44 24c0 2.633-.508 5.146-1.433 7.448c-.936 2.331-4.129.071-7.346 3.521c-3.216 3.45-.71 6.267-3.204 7.36A19.931 19.931 0 0 1 24 44C12.954 44 4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Z"/><path stroke-linecap="round" d="M20 25s.21 1.21 1 2s2 1 2 1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEgg0)"/>`),
		g.Group(children),
	)
}