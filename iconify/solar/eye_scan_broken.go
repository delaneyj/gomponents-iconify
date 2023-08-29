package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeScanBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 15c0 3.771 0 4.657-1.172 5.828C19.657 22 17.771 22 14 22M10 2C6.229 2 4.343 2 3.172 3.172C2 4.343 2 5.229 2 9m10-2C9.073 7 7.08 8.562 5.892 9.94C5.297 10.63 5 10.975 5 12c0 1.025.297 1.37.892 2.06C7.08 15.438 9.072 17 12 17c2.927 0 4.92-1.562 6.108-2.94c.595-.69.892-1.035.892-2.06c0-1.025-.297-1.37-.892-2.06A9.067 9.067 0 0 0 16 8.125"/><circle cx="12" cy="12" r="2"/><path stroke-linecap="round" d="M10 22H9m-7-7c0 3.771 0 4.657 1.172 5.828c.653.654 1.528.943 2.828 1.07M14 2h1m7 7c0-3.771 0-4.657-1.172-5.828c-.653-.654-1.528-.943-2.828-1.07"/></g>`),
		g.Group(children),
	)
}