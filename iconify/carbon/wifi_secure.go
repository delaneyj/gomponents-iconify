package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiSecure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 20v-3a4 4 0 0 0-8 0v3a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-6-3a2 2 0 0 1 4 0v3h-4zm-2 11v-6h8v6zM6.105 7.105l1.414 1.414a11.98 11.98 0 0 1 16.962 0l1.414-1.414a13.978 13.978 0 0 0-19.79 0z"/><path fill="currentColor" d="M16 8a8.97 8.97 0 0 0-6.36 2.64l1.412 1.413a6.989 6.989 0 0 1 9.895 0l1.414-1.414A8.972 8.972 0 0 0 16 8Z"/>`),
		g.Group(children),
	)
}