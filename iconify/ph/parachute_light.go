package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParachuteLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M230 120a102 102 0 0 0-204 0a6 6 0 0 0 2.27 4.69l.13.11L122 195v23h-10a6 6 0 0 0 0 12h32a6 6 0 0 0 0-12h-10v-23l93.6-70.2a6 6 0 0 0 2.4-4.8Zm-12.2-6h-43.9c-1.3-42.92-16.5-68.62-28.43-82.3A90.2 90.2 0 0 1 217.8 114ZM128 31.43A77.14 77.14 0 0 1 143.42 49c13.84 21.08 17.82 46 18.47 65H94.11c1.06-31.88 10.49-52.86 18.47-65A76.69 76.69 0 0 1 128 31.43ZM157.8 126L128 179.65L98.2 126Zm-73.33 0l24.62 44.32L50 126Zm87.06 0H206l-59.09 44.32Zm-61-94.3C98.6 45.38 83.4 71.08 82.1 114H38.2a90.2 90.2 0 0 1 72.33-82.3Z"/>`),
		g.Group(children),
	)
}