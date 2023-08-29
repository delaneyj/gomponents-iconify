package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdFlashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M459.9 435.5L76.1 52.5 51.9 76.6 160 184.3V272h64v192l72-144 139.9 139.5z" fill="currentColor"/><path d="M352 208h-64l64-160H160v40.3l168 167.6z" fill="currentColor"/>`),
		g.Group(children),
	)
}