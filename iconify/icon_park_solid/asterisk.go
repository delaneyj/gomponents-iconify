package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAsterisk0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" d="M15 24h18m-13.5-7.794l9 15.588m0-15.588l-9 15.588"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAsterisk0)"/>`),
		g.Group(children),
	)
}