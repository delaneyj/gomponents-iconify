package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoshare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m44.345 20.827l-3.55-1.02v.079a17.178 17.178 0 0 0-2.017-4.86l.063.062l1.789-3.23a1.595 1.595 0 0 0-.268-1.9l-2.32-2.32a1.595 1.595 0 0 0-1.9-.268l-3.231 1.79l.06.06a17.175 17.175 0 0 0-4.86-2.015h.082l-1.02-3.55A1.595 1.595 0 0 0 25.64 2.5h-3.28a1.595 1.595 0 0 0-1.533 1.155l-1.02 3.55h.079a17.176 17.176 0 0 0-4.86 2.017l.063-.063l-3.232-1.789a1.595 1.595 0 0 0-1.9.268l-2.32 2.32a1.595 1.595 0 0 0-.267 1.9l1.79 3.231l.06-.06a17.171 17.171 0 0 0-2.015 4.86v-.082l-3.55 1.02A1.595 1.595 0 0 0 2.5 22.36v3.28a1.595 1.595 0 0 0 1.155 1.533l3.55 1.02v-.079a17.178 17.178 0 0 0 2.017 4.86l-.063-.062l-1.789 3.23a1.595 1.595 0 0 0 .268 1.9l2.32 2.32a1.595 1.595 0 0 0 1.9.268l3.231-1.79l-.06-.06a17.173 17.173 0 0 0 4.86 2.016h-.082l1.02 3.55A1.595 1.595 0 0 0 22.36 45.5h3.28a1.595 1.595 0 0 0 1.533-1.155l1.02-3.55h-.079a17.178 17.178 0 0 0 4.86-2.017l-.063.063l3.231 1.789a1.595 1.595 0 0 0 1.9-.267l2.32-2.32a1.595 1.595 0 0 0 .268-1.9l-1.79-3.233l-.06.06a17.17 17.17 0 0 0 2.015-4.86v.082l3.55-1.02A1.595 1.595 0 0 0 45.5 25.64v-3.28a1.595 1.595 0 0 0-1.155-1.533Z"/><circle cx="27.081" cy="19.245" r="2.471" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.081" cy="28.755" r="2.471" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.919" cy="23.813" r="2.471" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.025 25.088l3.947 2.39m-1.123-6.424l1.085-.608m-3.866 2.164l1.158-.648"/>`),
		g.Group(children),
	)
}