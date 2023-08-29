package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hyperrail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.83 4.5v39m24.165-39v39M10.734 8.672h26.532c1.475 0 2.67 1.116 2.67 2.492h0c0 1.376-1.195 2.492-2.67 2.492H10.734c-1.475 0-2.67-1.116-2.67-2.492h0c0-1.376 1.195-2.492 2.67-2.492Zm0 12.904h26.532c1.475 0 2.67 1.115 2.67 2.491h0c0 1.377-1.195 2.492-2.67 2.492H10.734c-1.475 0-2.67-1.115-2.67-2.492h0c0-1.376 1.195-2.491 2.67-2.491Zm0 12.764h26.532c1.475 0 2.67 1.115 2.67 2.491h0c0 1.377-1.195 2.492-2.67 2.492H10.734c-1.475 0-2.67-1.115-2.67-2.492h0c0-1.376 1.195-2.491 2.67-2.491Z"/>`),
		g.Group(children),
	)
}