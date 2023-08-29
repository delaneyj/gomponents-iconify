package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThreeWayTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.062 4.938L8.958.059L6.03 4.94H8v1.806c0 .74.234 1.161.974 1.161c.737 0 .995-.421.995-1.161V4.939l2.093-.001zm-1.021 9.996l5.782-.006l-2.552-5.087l-1.055 1.662l-1.525-.968c-.625-.397-1.105-.424-1.502.2c-.396.622-.178 1.065.447 1.462l1.525.968l-1.12 1.769zm-4.102-.084l-5.782-.006l2.552-5.087l1.055 1.662l1.525-.968c.625-.397 1.106-.424 1.502.2c.395.622.178 1.065-.447 1.462l-1.525.968l1.12 1.769z"/>`),
		g.Group(children),
	)
}