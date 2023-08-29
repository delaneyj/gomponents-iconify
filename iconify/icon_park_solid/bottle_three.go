package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBottleThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.188 10h5.625L33 21.18V44H15V21.18L21.188 10Z"/><path fill="#fff" d="M20 4h8v6h-8z"/><rect width="6" height="12" x="21" y="23" fill="#fff" rx="3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBottleThree0)"/>`),
		g.Group(children),
	)
}