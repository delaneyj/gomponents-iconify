package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8-2a2 2 0 1 0 4 0a2 2 0 0 0-4 0ZM5.25 4A3.25 3.25 0 0 0 2 7.25v17.5A3.25 3.25 0 0 0 5.25 28h21.5A3.25 3.25 0 0 0 30 24.75V7.25A3.25 3.25 0 0 0 26.75 4H5.25ZM15 16a4 4 0 1 1-4-4h10a4 4 0 1 1-3.465 2h-3.07c.34.588.535 1.271.535 2Z"/>`),
		g.Group(children),
	)
}