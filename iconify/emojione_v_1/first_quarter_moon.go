package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M32 64C14.396 64 .125 49.73.125 32.125S14.396.249 32 .249"/><path fill="#f5eb35" d="M32.07 64c17.604 0 31.875-14.27 31.875-31.875C63.945 14.521 49.674.25 32.07.25"/><circle cx="19.334" cy="49.03" r="9.226" fill="#4f6977"/><path fill="#e0cf35" d="M41.905 24.488a3.919 3.919 0 0 1-7.836 0a3.917 3.917 0 0 1 3.915-3.916a3.92 3.92 0 0 1 3.921 3.916m8.782 23.882a3.84 3.84 0 0 1-7.678 0a3.84 3.84 0 0 1 7.678 0"/><g fill="#4f6977"><circle cx="5.967" cy="36.542" r="3.845"/><circle cx="6.313" cy="18.918" r="2.195"/></g><circle cx="56.23" cy="32.896" r="2.195" fill="#e0cf35"/><path fill="#4f6977" d="M20.968 19.656a3.433 3.433 0 1 1-6.867.001a3.433 3.433 0 0 1 6.867-.001"/><path fill="#e0cf35" d="M47.731 11.08a4.835 4.835 0 1 1-9.669-.005a4.835 4.835 0 0 1 9.669.005"/>`),
		g.Group(children),
	)
}