package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.538 34.577C20.27 36.114 14.61 25.293 15.63 24.228c.67-.7 1.689-.746 2.587-.76c4.109-.123 9.907 4.686 11.094 4.656c2.283-.046 5.707-3.683 7.96-6.681c.943-1.248 6.833-3.135 6.178.152C41.32 32.28 34.12 34.38 32.553 34.577h-.015Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.478 13.418c12.25-1.507 17.912 9.298 16.892 10.364c-.67.715-1.689.76-2.587.76c-4.109.122-9.907-4.687-11.094-4.656c-2.283.045-5.707 3.683-7.96 6.68c-.943 1.249-6.833 3.151-6.178-.151c2.145-10.7 9.344-12.8 10.927-12.982v-.015Z"/>`),
		g.Group(children),
	)
}