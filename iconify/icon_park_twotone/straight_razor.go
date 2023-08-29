package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightRazor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStraightRazor0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="6" x="3.609" y="36.534" fill="#555" rx="2" transform="rotate(-10 3.61 36.534)"/><path d="m44 40l-4-4"/><path fill="#555" d="m8 4l18.385 18.385l-4.243 4.242L9.414 13.9c-2.828-2.83-2.828-4.244-2.828-5.658C6.586 6.828 8 4 8 4Z"/><path d="m8 4l18 18l9 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStraightRazor0)"/>`),
		g.Group(children),
	)
}