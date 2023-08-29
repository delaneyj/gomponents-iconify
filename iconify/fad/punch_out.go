package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PunchOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M233.172 67.494a3.999 3.999 0 0 0-4.006-3.994h-55.494c-8.837 0-18.43 6.728-21.436 15.054l-54.837 82.392C94.397 169.26 84.8 176 75.971 176H30.765a4.002 4.002 0 0 0-4.006 4.007v8.986A4.004 4.004 0 0 0 30.757 193h58.708c8.835 0 18.381-6.752 21.324-15.089l54.057-82.822C167.787 86.756 177.337 80 186.171 80h43c2.21 0 4-1.797 4-3.994v-8.512z"/>`),
		g.Group(children),
	)
}