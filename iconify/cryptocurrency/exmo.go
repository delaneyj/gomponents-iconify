package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm3.7-18.945L19.171 12l-1.1.5l-2.885 7.797l1.1-.5l.526 1.055l.018-.047l2.868-7.75zm7.18.183l.064-.183L26.419 12l-1.1.5l-2.867 7.76l-.067.182l1.1-.5l.527 1.058l2.868-7.762zm-6.14 6.712l1.689-4.563l-1.103.5l-.524-1.057l-1.694 4.562l.525 1.058l1.107-.5zm-9.137-4.5H6.558l.86.8l-.86.813h5.04l.856-.813l-.851-.8zM5.86 18.833L5 19.64l.86.805h8.155l-.857-.805l.857-.808H5.86zm2.501-6.768l-.86.808l.86.805h8.15l-.854-.806l.855-.807h-8.15z"/>`),
		g.Group(children),
	)
}