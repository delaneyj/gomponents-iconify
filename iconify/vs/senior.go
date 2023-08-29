package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Senior(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M642 0q-78 0-132.5 55T455 188t54.5 132.5T642 375q79 0 133.5-54.5T830 188T775.5 55T642 0zm-92 428q-24-52-88-56q-24-2-94-2q-159 5-286 205q-61 97-74.5 201T38 980l62 269l-75 406q-9 52 21.5 94t80.5 43q119 0 141-89l89-389q8-38 8-53t-9-56l-42-248l110-292q9 13 32.5 50t39.5 60t40 52.5t48.5 49.5t49.5 31q23 10 328 74q41 8 68.5-24t27.5-66q0-43-57-74l-288-82zm241 767q0-71 50.5-121t122.5-50q70 0 120.5 50t50.5 121l-2 554q0 18-12.5 30.5T1090 1792q-17 0-30-12.5t-13-30.5v-554q0-36-24-61t-59-25q-36 0-60.5 25t-24.5 61q0 18-13 30t-30 12q-18 0-31.5-12t-13.5-30z"/>`),
		g.Group(children),
	)
}