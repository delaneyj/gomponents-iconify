package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spacenow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.94 28.054l-.073.04a21.506 21.506 0 0 0 3.072-2.493c6.619-6.619 10.016-12.576 9.512-20.052c-7.476-.505-13.433 2.893-20.051 9.512a21.501 21.501 0 0 0-2.495 3.073l.042-.074C15.552 16.823 8.765 19.164 5.5 23.135l11.513.108l.026-.047a60.703 60.703 0 0 0-2.107 5.024l4.848 4.848c1.612-.595 3.33-1.3 5.023-2.107l-.046.026l.108 11.513c3.97-3.265 6.312-10.052 5.075-14.446Zm-13.503 7.519a13.365 13.365 0 0 1-6.903 2.893a13.365 13.365 0 0 1 2.893-6.903l.99 3.02Z"/><circle cx="31.596" cy="16.404" r="2.434" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}