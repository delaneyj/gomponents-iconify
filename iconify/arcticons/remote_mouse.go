package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteMouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="21.522" height="34.878" x="10.222" y="9.578" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.761" transform="rotate(45 20.983 27.017)"/><circle cx="25.61" cy="22.39" r="1.92" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.724 14.577a9.661 9.661 0 0 0-5.257-5.283m9.033 3.678A13.777 13.777 0 0 0 35.028 5.5"/>`),
		g.Group(children),
	)
}