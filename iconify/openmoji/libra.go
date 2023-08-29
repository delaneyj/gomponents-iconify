package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12h48v48H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="2" d="M12 12h48v48H12z"/><path stroke-width="3" d="M30.53 39.5H20m32 0h-9.53m9.53 5H20m10.53-5a8.443 8.443 0 0 1-1.83-9.201a8.443 8.443 0 0 1 7.8-5.212a8.443 8.443 0 0 1 7.8 5.212a8.443 8.443 0 0 1-1.83 9.201"/></g>`),
		g.Group(children),
	)
}