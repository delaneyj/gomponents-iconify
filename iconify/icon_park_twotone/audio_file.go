package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAudioFile0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 44V4h23l9 10.5V44H8Z"/><path d="m32 14l-6 2.969V31.5"/><circle cx="20.5" cy="31.5" r="5.5" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAudioFile0)"/>`),
		g.Group(children),
	)
}