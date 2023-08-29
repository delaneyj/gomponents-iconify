package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phototan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.17 38.39h4.11v4.11h-4.11Zm-16.45 0h4.11v4.11h-4.11Zm8.22-8.22h4.12v4.11h-4.12Zm-8.22 0h4.11v4.11h-4.11Zm16.45-8.23h4.11v4.12h-4.11Zm-16.45 0h4.11v4.12h-4.11Zm-8.22 0h4.11v4.12H5.5Zm32.89 0h4.11v4.12h-4.11Zm0-8.22h4.11v4.11h-4.11Zm-16.45 0h4.12v4.11h-4.12Zm8.23-8.22h4.11v4.11h-4.11Zm-8.23 0h4.12v4.11h-4.12Zm-8.22 0h4.11v4.11h-4.11Zm-8.22 0h4.11v4.11H5.5Z"/>`),
		g.Group(children),
	)
}