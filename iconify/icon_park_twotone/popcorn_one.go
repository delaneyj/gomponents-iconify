package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopcornOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPopcornOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 16h34l-7 28H14L7 16Z"/><path d="M20 16v28m8-28v28"/><path fill="#555" d="M33 9a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 33 9Zm-9 0a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 24 9Zm-9 0a4 4 0 0 0-2.646 7h5.292A4 4 0 0 0 15 9Z"/><path d="M22.874 9a4 4 0 1 0-7.748 0m17.748 0a4 4 0 1 0-7.748 0M16 16h16M16 44h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPopcornOne0)"/>`),
		g.Group(children),
	)
}