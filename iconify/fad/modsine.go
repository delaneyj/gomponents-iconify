package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modsine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M48 128c-1.955-29.248 19.364-64 37.364-64c18 0 36.136 13.843 36.136 64.5s19.136 80.5 49.136 80.5c30 0 53.364-40.125 53.364-80.5c-8.182 0-7.273-.752-16 0c0 32.35-20.455 64.45-37.364 64.45s-33.909-13.542-33.909-64.45S120.273 48 85.364 48C50.454 48 32 88.626 32 127.748c6 0 8.364.252 16 .252z"/>`),
		g.Group(children),
	)
}