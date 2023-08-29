package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceInfoPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.95 19.05c-2.734-2.734-15.84-5.941-15.84-5.941s3.207 13.107 5.94 15.841a7 7 0 0 0 9.9-9.9Z"/>`),
		g.Group(children),
	)
}