package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLayersOutline0"><g id="evaLayersOutline1"><path id="evaLayersOutline2" fill="currentColor" d="M21 11.35a1 1 0 0 0-.61-.86l-2.15-.92l2.26-1.3a1 1 0 0 0 .5-.92a1 1 0 0 0-.61-.86l-8-3.41a1 1 0 0 0-.78 0l-8 3.41a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l2.26 1.3l-2.15.92a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l2.26 1.3l-2.15.92a1 1 0 0 0-.61.86a1 1 0 0 0 .5.92l8 4.6a1 1 0 0 0 1 0l8-4.6a1 1 0 0 0 .5-.92a1 1 0 0 0-.61-.86l-2.15-.92l2.26-1.3a1 1 0 0 0 .5-.92Zm-9-6.26l5.76 2.45L12 10.85L6.24 7.54Zm-.5 7.78a1 1 0 0 0 1 0l3.57-2l1.69.72L12 14.85l-5.76-3.31l1.69-.72Zm6.26 2.67L12 18.85l-5.76-3.31l1.69-.72l3.57 2.05a1 1 0 0 0 1 0l3.57-2.05Z"/></g></g>`),
		g.Group(children),
	)
}