package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 16c0-7.109 10.667-7.109 10.667 0S16 23.109 16 16zM5.333 26.667a5.33 5.33 0 0 1 5.333-5.333h5.333v5.333c0 7.109-10.667 7.109-10.667 0zM16 0v10.667h5.333c7.109 0 7.109-10.667 0-10.667zM5.333 5.333a5.33 5.33 0 0 0 5.333 5.333h5.333V-.001h-5.333a5.33 5.33 0 0 0-5.333 5.333zm0 10.667a5.33 5.33 0 0 0 5.333 5.333h5.333V10.666h-5.333a5.33 5.33 0 0 0-5.333 5.333z"/>`),
		g.Group(children),
	)
}