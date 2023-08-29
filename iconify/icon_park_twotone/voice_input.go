package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVoiceInput0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 30c0 6.627 5.373 12 12 12s10-4 11-7l1.538-5L35 9l9 33m-3.273-12H28.54"/><path fill="#555" d="M22 15a6 6 0 0 0-12 0v15a6 6 0 0 0 12 0V15Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVoiceInput0)"/>`),
		g.Group(children),
	)
}