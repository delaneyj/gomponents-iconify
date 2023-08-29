package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrophy0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M24 30c6.627 0 12-5.53 12-12.353V4H12v13.647C12 24.47 17.373 30 24 30Z"/><path stroke-linecap="round" d="M12 21V11H4c0 6.667 4 10 8 10Zm24 0V11h8c0 6.667-4 10-8 10Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M24 32v4"/><path fill="#fff" d="m15 42l3.69-6h10.352L33 42H15Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrophy0)"/>`),
		g.Group(children),
	)
}