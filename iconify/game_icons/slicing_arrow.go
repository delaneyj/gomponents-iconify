package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlicingArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m94.027 17.81l-.763 71.667l-71.254.355l85.51 64.35l27.94-9.082l277.542 277.537l-39.828 9.345l118.6 56.215l-56.22-118.597l-9.34 39.818l-277.376-277.37l9.54-28.726l-64.35-85.51zm83.41.465L393.99 308.74L280.887 18.275h-103.45zm123.506 0l129.922 333.66l-39.35-333.66h-90.572zm109.39 0l34.173 289.768l51.24-289.768h-85.414zM20.98 174.443V277.9l295.903 117.153L20.98 174.443zm0 123.557v90.557l335.89 42.425L20.98 298zm0 109.393v85.36l281.59-49.792l-281.59-35.567z"/>`),
		g.Group(children),
	)
}