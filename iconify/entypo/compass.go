package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.454 14.548s4.568-.627 6.518-2.576s2.576-6.518 2.576-6.518s-4.569.627-6.518 2.576s-2.576 6.518-2.576 6.518zm3.563-5.533c.818-.818 2.385-1.4 3.729-1.762c-.361 1.342-.945 2.92-1.76 3.732a1.39 1.39 0 0 1-1.969 0a1.391 1.391 0 0 1 0-1.97zM10.001.4C4.698.4.4 4.698.4 10a9.6 9.6 0 0 0 9.601 9.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zM10 17.6a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2z"/>`),
		g.Group(children),
	)
}