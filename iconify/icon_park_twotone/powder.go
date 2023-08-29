package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPowder0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M44 40V20a3 3 0 0 0-3-3h-5.285c-1.541 0-2.892 1.243-3.924 2.388C30.633 20.674 28.377 22 24 22c-4.377 0-6.633-1.326-7.791-2.612C15.177 18.243 13.826 17 12.285 17H7a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 17c0 2.761-4.03 5-9 5s-9-2.239-9-5"/><ellipse cx="24" cy="10" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="9" ry="5"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 17v-7m-18 7v-7"/><ellipse cx="27" cy="10" fill="#fff" rx="2" ry="1"/><ellipse cx="21" cy="10" fill="#fff" rx="2" ry="1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPowder0)"/>`),
		g.Group(children),
	)
}