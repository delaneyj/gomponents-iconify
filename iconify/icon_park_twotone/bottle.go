package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBottle0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M36 18H12v26h24V18Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M19.894 11h8.212a4 4 0 0 1 3.8 2.75L33.303 18H14.695l1.4-4.25a4 4 0 0 1 3.8-2.75Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M21 11V7a3 3 0 1 1 6 0v4m-8.5 15v10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBottle0)"/>`),
		g.Group(children),
	)
}