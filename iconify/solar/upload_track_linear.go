package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTrackLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 15a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 0V9"/><path stroke-linecap="round" d="m14.058 11.03l-1.316-.659a2.313 2.313 0 0 1-.35-.194a1 1 0 0 1-.374-.607C12 9.477 12 9.375 12 9.171c0-.486 0-.728.06-.893a1 1 0 0 1 1.056-.653c.174.02.391.129.826.346l1.316.658c.183.092.274.137.35.195a1 1 0 0 1 .374.606c.018.093.018.195.018.4c0 .485 0 .728-.06.893a1 1 0 0 1-1.056.652c-.174-.02-.391-.129-.826-.346Z"/><path stroke-linecap="round" d="M14 21.8c-.646.131-1.315.2-2 .2c-5.523 0-10-4.477-10-10S6.477 2 12 2s10 4.477 10 10c0 .685-.069 1.354-.2 2"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 22v-7m0 0l2.5 2.5M18 15l-2.5 2.5"/></g>`),
		g.Group(children),
	)
}