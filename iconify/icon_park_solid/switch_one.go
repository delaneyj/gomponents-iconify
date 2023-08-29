package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSwitchOne0"><g fill="none" stroke-width="4"><circle cx="24" cy="24" r="19" fill="#fff" stroke="#fff"/><rect width="10.519" height="24.012" x="36.006" y="19.334" fill="#000" stroke="#000" rx="5.259" transform="rotate(90 36.006 19.334)"/><rect width="10" height="10" x="36.006" y="29.852" stroke="#000" rx="5" transform="rotate(-180 36.006 29.852)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSwitchOne0)"/>`),
		g.Group(children),
	)
}