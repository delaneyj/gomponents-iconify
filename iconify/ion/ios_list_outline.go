package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosListOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M432 80v352H80V80h352m16-16H64v384h384V64z" fill="currentColor"/><path d="M192 152h192v16H192z" fill="currentColor"/><path d="M192 248h192v16H192z" fill="currentColor"/><path d="M192 344h192v16H192z" fill="currentColor"/><circle cx="144" cy="160" r="16" fill="currentColor"/><circle cx="144" cy="256" r="16" fill="currentColor"/><circle cx="144" cy="352" r="16" fill="currentColor"/>`),
		g.Group(children),
	)
}