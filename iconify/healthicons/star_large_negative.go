package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarLargeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsStarLargeNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.903 7.064a1.005 1.005 0 0 0-1.806 0L18.315 16.8c-.147.298-.43.505-.759.553L6.864 18.914c-.826.121-1.156 1.141-.558 1.726l7.738 7.579c.237.232.346.567.29.895l-1.827 10.7c-.141.827.722 1.458 1.461 1.068l9.564-5.053c.294-.155.644-.155.938 0l9.564 5.053c.739.39 1.602-.24 1.461-1.067l-1.827-10.7a1.015 1.015 0 0 1 .29-.896l7.738-7.579c.597-.585.268-1.605-.558-1.726l-10.694-1.56a1.008 1.008 0 0 1-.758-.554l-4.782-9.736Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsStarLargeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}