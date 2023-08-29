package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTriangleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaAlertTriangleFill0"><g id="evaAlertTriangleFill1"><path id="evaAlertTriangleFill2" fill="currentColor" d="M22.56 16.3L14.89 3.58a3.43 3.43 0 0 0-5.78 0L1.44 16.3a3 3 0 0 0-.05 3A3.37 3.37 0 0 0 4.33 21h15.34a3.37 3.37 0 0 0 2.94-1.66a3 3 0 0 0-.05-3.04ZM12 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm1-4a1 1 0 0 1-2 0V9a1 1 0 0 1 2 0Z"/></g></g>`),
		g.Group(children),
	)
}