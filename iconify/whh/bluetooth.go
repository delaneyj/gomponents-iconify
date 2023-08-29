package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M750 335q-11 11-30 15L471 513l279 183q18 18 18 43.5T750 783q-13 13-35 16l-347 205q-3 1-10 5.5t-11.5 7.5t-11 5.5t-13.5 2.5q-28 2-48.5-18T257 956V653l-152 99q-18 18-43.5 18T18 752T0 708.5T18 665l232-152L18 361Q0 343 0 317.5T18 274t43.5-18t43.5 18l152 99V70q-4-31 16.5-51T322 1q5 0 10 1t10 4t8 4.5t9.5 6T368 23l382 225q18 18 18 43.5T750 335zM385 855l205-120l-205-134v254zm0-684v254l205-134z"/>`),
		g.Group(children),
	)
}