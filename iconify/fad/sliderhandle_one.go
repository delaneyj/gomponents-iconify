package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderhandleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M64.254 32.888C63.56 28.532 66.584 25 70.996 25h113.008c4.416 0 7.477 3.545 6.875 7.913c0 0-6.879 43.782-6.879 95.218c0 51.435 6.875 94.961 6.875 94.961c.621 4.368-2.46 7.908-6.871 7.908H70.996c-4.416 0-7.402-3.537-6.714-7.899c0 0 7.718-42.667 7.718-94.474c0-51.808-7.746-95.739-7.746-95.739zM80 40s8 44.261 8 88.521S80 217 80 217h96s-6.533-44.446-6.533-88.479C169.467 84.49 176 40 176 40H80z"/><path d="M95 57h65l-4 16H98zm5 125h55l5 19H96zm4-58h48v9h-48z"/></g>`),
		g.Group(children),
	)
}