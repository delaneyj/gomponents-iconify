package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoomin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M120.46 158.29c-30.166 8.65-61.631-8.792-70.281-38.957c-8.65-30.165 8.792-61.63 38.957-70.28c30.165-8.65 61.63 8.792 70.28 38.957c4.417 15.403-1.937 38.002-9.347 50.872c-.614 1.067 59.212 53.064 59.212 53.064l-17.427 17.63l-53.514-56.72s-10.233 3.241-17.88 5.434zM104 144c22.091 0 40-17.909 40-40s-17.909-40-40-40s-40 17.909-40 40s17.909 40 40 40z"/><path d="M111.912 80.047h-15.95v16.037H80v15.992h15.962V128h15.95v-15.924H128V96.084h-16.088z"/></g>`),
		g.Group(children),
	)
}