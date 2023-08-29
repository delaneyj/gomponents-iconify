package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M704 768q-159 0-271.5-112.5T320 384t112.5-271.5T704 0t271.5 112.5T1088 384T975.5 655.5T704 768zm960 128h352q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-352v352q0 13-9.5 22.5t-22.5 9.5h-192q-13 0-22.5-9.5t-9.5-22.5v-352h-352q-13 0-22.5-9.5t-9.5-22.5V928q0-13 9.5-22.5t22.5-9.5h352V544q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v352zm-736 224q0 52 38 90t90 38h256v238q-68 50-171 50H267q-121 0-194-69T0 1277q0-53 3.5-103.5t14-109T44 956t43-97.5t62-81t85.5-53.5T346 704q19 0 39 17q79 61 154.5 91.5T704 843t164.5-30.5T1023 721q20-17 39-17q132 0 217 96h-223q-52 0-90 38t-38 90v192z"/>`),
		g.Group(children),
	)
}