package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVoicemail0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M11 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm26 0a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke-linecap="round" d="M12 31h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVoicemail0)"/>`),
		g.Group(children),
	)
}