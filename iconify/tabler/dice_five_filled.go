package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFiveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005.195v12.666c0 1.96-1.537 3.56-3.472 3.662l-.195.005H5.667a3.667 3.667 0 0 1-3.662-3.472L2 18.333V5.667c0-1.96 1.537-3.56 3.472-3.662L5.667 2h12.666zM15.5 14a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm-7 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm3.5-3.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM8.5 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm7 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3z"/></g>`),
		g.Group(children),
	)
}