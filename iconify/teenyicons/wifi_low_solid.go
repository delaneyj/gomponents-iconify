package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLowSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.219 9.318c1.155-1.4 2.698-2.161 4.28-2.161c1.584 0 3.127.762 4.282 2.161l.771-.636C11.232 7.08 9.417 6.157 7.5 6.157c-1.918 0-3.732.924-5.052 2.525l.77.636ZM6 11.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}