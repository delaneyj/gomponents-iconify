package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zendframework(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M745 256H0q1-5 3-14.5t11.5-36T36 155t35-54t50.5-51t69-35.5T279 0h745q-1 5-3 14.5t-11.5 36T988 101t-35 53.5t-50.5 51t-69 36T745 256zm-466 64h425q-1 5-3 14.5t-11.5 36T668 421t-35 53.5t-50.5 51t-69 36T425 576H0q1-5 3-14.5t11.5-36T36 475t35-54t50.5-51t69-35.5T279 320zm0 320h169q-1 5-3 14.5t-11.5 36T412 741t-35 53.5t-50.5 51t-69 36T169 896H0q1-5 3-14.5t11.5-36T36 795t35-54t50.5-51t69-35.5T279 640z"/>`),
		g.Group(children),
	)
}