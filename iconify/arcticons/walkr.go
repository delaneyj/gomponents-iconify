package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walkr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.316 15.793a2.438 2.438 0 0 0 1.716.703h.002a2.448 2.448 0 0 0-.002-4.896h0a2.439 2.439 0 0 0-2.43 2.448a2.383 2.383 0 0 0 .714 1.745ZM6.838 7.51a.675.675 0 0 1 .206-.479a.807.807 0 0 1 .516-.197a23.863 23.863 0 0 1 9.126 1.51a24.875 24.875 0 0 1 7.626 5.487c1.153 1.154 2.391 2.533 3.751 4.146l8.057.422a.806.806 0 0 1 .553.347l4.765 8.16a.656.656 0 0 1-.103.825l-1.37 1.351a.676.676 0 0 1-.487.188h-.197l-5.824-1.801l-5.966 5.975l1.801 5.871a.685.685 0 0 1-.16.675l-1.37 1.36a.75.75 0 0 1-.477.197a.646.646 0 0 1-.347-.093l-8.16-4.756a.74.74 0 0 1-.338-.553l-.422-8.057q-2.42-2.017-4.146-3.752a24.386 24.386 0 0 1-5.43-7.653a24.057 24.057 0 0 1-1.604-9.202Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.883 6.964a21.498 21.498 0 1 1-3.887 3.876"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.757 2.663a5.595 5.595 0 1 0 10.337 4.284M14.363 43.224A7.32 7.32 0 0 0 4.45 32.982"/>`),
		g.Group(children),
	)
}