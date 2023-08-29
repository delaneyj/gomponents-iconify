package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWebComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" d="m179.9 388l-76.16-132l76.16 132zm0 0h152.21l76.15-132l-76.15-132H179.9l-76.16 132l76.16 132zm-76.16-132l76.16-132l-76.16 132z"/><path fill="currentColor" d="M496 256L376 48H239.74l-43.84 76h136.21l76.15 132l-76.15 132H195.9l43.84 76H376l120-208z"/><path fill="currentColor" d="m179.9 388l-76.16-132l76.16-132l43.84-76H136L16 256l120 208h87.74l-43.84-76z"/>`),
		g.Group(children),
	)
}