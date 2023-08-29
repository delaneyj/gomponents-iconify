package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openvegemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.502a13.09 13.09 0 0 0-13.09 13.09c0 10.25 10 22.61 12.61 25.63a.8.8 0 0 0 1.21 0c2.55-3 12.36-15.38 12.36-25.63A13.09 13.09 0 0 0 24 4.502Z"/><circle cx="24.004" cy="17.596" r="10.09" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.266 14.645a2.085 2.085 0 0 1 1.051-1.703c.994-.673 1.725-.423 2.054-.072a1.641 1.641 0 0 1 .091 1.966c-.73 1.061-1.72.673-2.42 1.665c-1.89 2.679-2.955 4.384-3.485 8.049v.001l-1.16-5s-.918-4.672-3.86-6.775"/>`),
		g.Group(children),
	)
}