package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h4A1.5 1.5 0 0 1 7 1.5v4A1.5 1.5 0 0 1 5.5 7h-4A1.5 1.5 0 0 1 0 5.5v-4ZM1.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4Zm6.5.5A1.5 1.5 0 0 1 9.5 0h4A1.5 1.5 0 0 1 15 1.5v4A1.5 1.5 0 0 1 13.5 7h-4A1.5 1.5 0 0 1 8 5.5v-4ZM9.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM4 4H3V3h1v1Zm8 0h-1V3h1v1ZM0 9.5A1.5 1.5 0 0 1 1.5 8h4A1.5 1.5 0 0 1 7 9.5v4A1.5 1.5 0 0 1 5.5 15h-4A1.5 1.5 0 0 1 0 13.5v-4ZM1.5 9a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM8 8h4v1H9v3H8V8Zm7 1h-1V8h1v1ZM4 12H3v-1h1v1Zm10 0h-3v-1h4v4h-1v-3Zm-6 2h4v1H8v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}