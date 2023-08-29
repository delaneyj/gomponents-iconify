package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicRecordTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 17.5c0 1.096.271 2.129.75 3.035v.715a.75.75 0 0 1-1.493.102l-.007-.102v-2.268a6.75 6.75 0 0 1-6.246-6.496L4 12.25v-.5a.75.75 0 0 1 1.493-.102l.007.102v.5a5.25 5.25 0 0 0 5.034 5.246l.216.004H11Zm.175-1.504A6.513 6.513 0 0 1 15 11.498V6a4 4 0 0 0-8 0v6a4 4 0 0 0 4.175 3.996ZM20 17.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm3 0a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-9.5 0a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/>`),
		g.Group(children),
	)
}