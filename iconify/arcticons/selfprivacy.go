package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selfprivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.8 6.257a1.919 1.919 0 0 0-1.919 1.919v6.513h15.794l6.498 7.821v11.195a1.263 1.263 0 0 1-1.264 1.266H8.88v4.853a1.919 1.919 0 0 0 1.919 1.919H40.58a1.919 1.919 0 0 0 1.919-1.919V8.175a1.919 1.919 0 0 0-1.919-1.918Zm-5.039 28.22l3.095 4.138Zm3.12.495H6.767A1.264 1.264 0 0 1 5.5 33.705V15.953a1.263 1.263 0 0 1 1.266-1.265h2.115"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.27 23.684a3.073 3.073 0 1 1 4.386 2.78l1.484 5.599h-5.48l1.42-5.577a3.074 3.074 0 0 1-1.81-2.802"/>`),
		g.Group(children),
	)
}