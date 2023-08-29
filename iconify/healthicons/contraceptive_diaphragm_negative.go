package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContraceptiveDiaphragmNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsContraceptiveDiaphragmNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM37.955 25.13c.03-.345.045-.707.045-1.087C38 16.287 31.732 10 24 10s-14 6.287-14 14.043c0 .404.017.788.05 1.152l-.02.004c-.02.336-.03.675-.03 1.017c0 1.174.122 2.16.34 2.99a1 1 0 0 1-1.934.508C8.136 28.686 8 27.525 8 26.216c0-.169.002-.337.007-.505C3.199 27.67 1.429 37 24.27 37c22.905 0 20.135-8.89 15.723-11.267c.004.16.006.322.006.483c0 1.57-.196 2.93-.585 4.1a1 1 0 0 1-1.898-.632c.307-.921.483-2.06.483-3.468c0-.364-.012-.726-.035-1.084l-.01-.002Zm-16.66-9.1a1 1 0 0 1-.669 1.246A6.5 6.5 0 0 0 16 23.445a1 1 0 0 1-2-.017a8.5 8.5 0 0 1 6.05-8.067a1 1 0 0 1 1.246.67Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsContraceptiveDiaphragmNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}