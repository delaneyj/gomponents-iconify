package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClearFormat0"><g fill="none" stroke="#fff"><path fill="#fff" stroke-linejoin="round" stroke-width="4.302" d="M44.782 24.17L31.918 7.1L14.135 20.5L27.5 37l3.356-2.336L44.782 24.17Z"/><path stroke-linejoin="round" stroke-width="4.302" d="m27.5 37l-3.839 3.075l-10.563-.001l-2.6-3.45l-6.433-8.536L14.5 20.225"/><path stroke-linecap="round" stroke-width="4.5" d="M13.206 40.072h31.36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClearFormat0)"/>`),
		g.Group(children),
	)
}