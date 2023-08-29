package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 461L920 973q-3 21-25 35.5t-49 14.5H188q-27 0-49-14.5T113 973L2 461q-5-26 6.5-47T43 383q62 27 190.5 45.5T514 447q151 0 278.5-18T982 383q23 10 34.5 31t7.5 47zM833 95l128-32v264q-49 18-128 31V95zM642 0l127 64v304q-61 7-127 11V0zM450 382V32L578 0v382q-38 1-64.5 1t-63.5-1zM258 32L386 0v379q-67-3-128-11V32zM66 127l128-64v295q-79-13-128-30V127z"/>`),
		g.Group(children),
	)
}