package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M55.423.322H8.637C4.053.322.323 4.052.323 8.637v46.781c0 4.586 3.729 8.317 8.314 8.317h46.786c4.584 0 8.313-3.73 8.313-8.317V8.637c0-4.585-3.729-8.315-8.313-8.315zM51.52 38.528H38.528v12.991H25.539V38.528H12.543V25.529h12.996V12.538h12.989v12.991H51.52v12.999z"/>`),
		g.Group(children),
	)
}