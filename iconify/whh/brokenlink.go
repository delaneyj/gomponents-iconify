package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brokenlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M948.21 948.5q-76.5 76.5-185 76.5t-184.5-77l-52-51l113-64l25 24q39 40 95.5 40t96.5-40t40-96.5t-40-95.5l-188-189l76-101l203 204q77 76 77 184.5t-76.5 185zM231.71 392l345 121l-32 128l-350-130q-105-27-159-120t-26-197t122-157.5t199-26.5l214 87l-32 128l-211-94q-54-14-103 14t-64 82t13.5 102.5t83.5 62.5z"/>`),
		g.Group(children),
	)
}