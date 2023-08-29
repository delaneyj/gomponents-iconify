package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LipstickOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLipstickOne0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m39.435 5.323l1.414 1.414a4 4 0 0 1 0 5.657l-7.071 7.07l-4.243-4.242l9.9-9.9Zm-11.314 8.485l7.071 7.071l-11.314 11.314l-7.07-7.071z"/><path d="m15.394 23.707l9.9 9.9l-9.9 9.899l-9.9-9.9z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLipstickOne0)"/>`),
		g.Group(children),
	)
}