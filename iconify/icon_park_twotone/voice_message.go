package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVoiceMessage0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#fff" d="M17 25.9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21.95 28.85A6.978 6.978 0 0 0 24 23.9a6.978 6.978 0 0 0-2.05-4.95m4.95 14.849a13.956 13.956 0 0 0 4.1-9.9c0-3.866-1.567-7.366-4.1-9.899"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVoiceMessage0)"/>`),
		g.Group(children),
	)
}