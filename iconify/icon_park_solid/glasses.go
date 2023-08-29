package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGlasses0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="12" cy="35" r="7" fill="#fff"/><circle cx="36" cy="35" r="7" fill="#fff"/><path d="M5 34V10.883c0-1.391 0-2.087.378-2.61c.377-.525 1.037-.745 2.357-1.185L11 6m32 28V10.883c0-1.391 0-2.087-.377-2.61c-.378-.525-1.038-.745-2.358-1.185L37 6m-8 28a5 5 0 0 0-10 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGlasses0)"/>`),
		g.Group(children),
	)
}