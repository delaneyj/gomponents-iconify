package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M29 6v8h13v2H29v7h13V8c0-1.105-.836-2-1.867-2H29Zm0 19h13v7H29v-7Zm0 9h13v6c0 1.105-.836 2-1.867 2H29v-8Zm-2 0v8H15.867C14.836 42 14 41.105 14 40v-6h13Zm0-20H14V8c0-1.105.836-2 1.867-2H27v8Zm-3.948 2v7H27v-7h-3.948Zm0 9v7H27v-7h-3.948ZM6 17a1 1 0 0 1 1-1h13.158a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V17Zm3.607 2h2.26l1.834 3.754L15.64 19h2.112l-2.91 5l2.976 5H15.59l-1.999-3.93l-1.99 3.93H9.34l3.024-5.018L9.607 19Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}