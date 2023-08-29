package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditMutuel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.423 7.976L5.5 32.092l4.577 7.927h27.846l4.577-7.927L28.577 7.976h-9.154Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.864 17.606a6.94 6.94 0 0 1 3.758-1.092c2.818 0 5.25 1.656 6.388 4.052m-13.434 2.736c-.012.082-.012.176-.012.27c0 3.887 3.16 7.046 7.058 7.046c.106 0 .2 0 .294-.012a7.041 7.041 0 0 0 6.095-4.04c.035-.07.058-.129.093-.2m13.329-3.041c.012.082.012.164.012.247c0 3.734-2.924 6.787-6.6 7.022"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.133 17.594a7.044 7.044 0 0 0-10.123 2.971a6.766 6.766 0 0 0-.68 3.007c0 .998.21 1.95.586 2.806c.024.07.047.13.094.188m-6.094 4.04a7.043 7.043 0 0 0 4.016 9.418"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.829 40.024a7.043 7.043 0 0 0 4.603-6.611c0-.999-.211-1.962-.587-2.819a7.046 7.046 0 0 0-6.47-4.24c-.094 0-.188 0-.27.013c-.071 0-.13 0-.188.011"/>`),
		g.Group(children),
	)
}