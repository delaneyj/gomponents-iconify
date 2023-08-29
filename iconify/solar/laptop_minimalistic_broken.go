package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMinimalisticBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M3 10v4c0 1.886 0 2.828.586 3.414C4.172 18 5.114 18 7 18h10c1.886 0 2.828 0 3.414-.586C21 16.828 21 15.886 21 14V9c0-2.828 0-4.243-.879-5.121C19.243 3 17.828 3 15 3H9c-2.828 0-4.243 0-5.121.879c-.49.49-.707 1.146-.803 2.121M22 21h-6M2 21h10m3-6H9"/>`),
		g.Group(children),
	)
}