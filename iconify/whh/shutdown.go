package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shutdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm0 640q-26 0-45-18.5T448 704V320q0-26 19-45t45-19t45 19t19 45v384q0 27-19 45.5T512 768z"/>`),
		g.Group(children),
	)
}