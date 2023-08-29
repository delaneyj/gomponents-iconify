package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelLogoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsExcelLogoNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM29 14V6h11.133C41.164 6 42 6.895 42 8v15H29v-7h13v-2H29Zm13 11H29v7h13v2H29v8h11.133C41.164 42 42 41.105 42 40V25ZM27 42v-8H14v6c0 1.105.836 2 1.867 2H27ZM14 14V8c0-1.105.836-2 1.867-2H27v8H14Zm9.052 9v-7H27v7h-3.948Zm0 9v-7H27v7h-3.948ZM7 16a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V17a1 1 0 0 0-1-1H7Zm4.867 3h-2.26l2.758 4.982L9.34 29h2.26l1.991-3.93l2 3.93h2.226l-2.975-5l2.91-5h-2.113l-1.938 3.754L11.867 19Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsExcelLogoNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}