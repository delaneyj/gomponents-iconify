package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M460 187q0-76-67.5-129.5T230 4T67.5 57.5T0 187q0 65 51.5 115.5T181 367q-4 51-21 97q63-16 126-99q76-15 125-64.5T460 187zM329 300q-6 2-10 2h-18q59-43 59-115q0-79-70-121q6-2 19-2q46 0 78.5 34.5T420 183q0 44-26 77t-65 40z"/>`),
		g.Group(children),
	)
}