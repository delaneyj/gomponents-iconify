package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMute0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#000" stroke="#000" d="M29 14v20c-3.85 0-6.71-5.09-6.71-5.09H18.1c-.608 0-1.1-.497-1.1-1.11v-7.683c0-.614.492-1.111 1.1-1.111h4.19S25.15 14 29 14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMute0)"/>`),
		g.Group(children),
	)
}