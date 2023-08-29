package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModulzLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.26 3.167L4.37 5.333V1l2.89 2.167ZM1 8.222l2.889-2.166L1 3.889v4.333ZM1 14l2.889-2.167L1 9.667V14Zm6.74-5.778l2.89-2.166l-2.89-2.167v4.333ZM14 3.167l-2.889 2.166V1L14 3.167Zm-2.889 7.944L14 8.944l-2.889-2.166v4.333Zm-7.222 0L1 8.944l2.889-2.166v4.333Zm.481-5.055l2.89 2.166V3.89L4.37 6.056Zm-.481-.723L1 3.167L3.889 1v4.333ZM7.74 3.167l2.889 2.166V1L7.74 3.167ZM14 8.222l-2.889-2.166L14 3.889v4.333Zm-2.889 3.611L14 14V9.667l-2.889 2.166Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}