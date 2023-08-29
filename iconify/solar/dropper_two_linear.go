package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropperTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 10h-3"/><path d="M18 12.917c0 1.15-.895 2.083-2 2.083s-2-.933-2-2.083c0-.72.783-1.681 1.37-2.3a.862.862 0 0 1 1.26 0c.587.619 1.37 1.58 1.37 2.3Z"/><path stroke-linecap="round" d="M22 13h-2m-4 8v1M2 22V8c0-2.828 0-4.243.879-5.121C3.757 2 5.172 2 8 2h3.935a3 3 0 0 1 1.883.682l.034.028l.079.065a2.994 2.994 0 0 0 2.069.66M18.5 2l-.523.523A2.993 2.993 0 0 1 16 3.436M16 2v1.436m0 0V6"/><path stroke-linecap="round" d="M22 16.507V10c0-1.886 0-2.828-.586-3.414C20.828 6 19.886 6 18 6h-4c-1.886 0-2.828 0-3.414.586C10 7.172 10 8.114 10 10v6.507c0 .657.277 1.284.763 1.726a7.78 7.78 0 0 0 10.474 0A2.334 2.334 0 0 0 22 16.507Z"/></g>`),
		g.Group(children),
	)
}