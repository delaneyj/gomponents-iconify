package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PumpkinScary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l1.5 1l1.5-1l1.5 1l1.5-1m-5-4h.01M14 11h.01"/><path d="M17 6.082c2.609.588 3.627 4.162 2.723 7.983c-.903 3.82-2.75 6.44-5.359 5.853a3.355 3.355 0 0 1-.774-.279A3.728 3.728 0 0 1 12 20c-.556 0-1.09-.127-1.59-.362a3.296 3.296 0 0 1-.774.28c-2.609.588-4.456-2.033-5.36-5.853c-.903-3.82.115-7.395 2.724-7.983c1.085-.244 1.575.066 2.585.787C10.301 6.315 11.125 6 12 6c.876 0 1.699.315 2.415.87c1.01-.722 1.5-1.032 2.585-.788z"/><path d="M12 6c0-1.226.693-2.346 1.789-2.894L14 3"/></g>`),
		g.Group(children),
	)
}