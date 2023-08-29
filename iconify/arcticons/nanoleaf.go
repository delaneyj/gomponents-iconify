package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nanoleaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.541 40.569l-14.05-8.23a3.452 3.452 0 0 1-1.707-2.978V11.934c0-4.42 4.146-7.669 8.438-6.611l18.89 4.655a3.452 3.452 0 0 1 2.55 2.63l3.342 15.61a4.802 4.802 0 0 1-1.951 4.946l-10.34 7.202a4.802 4.802 0 0 1-5.172.203Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.792 36.828c1.883 2.232 2.65 5.294 2.65 5.294m-2.784-19.833l-2.596-2.791l-3.3.467l-.704 3.257l2.596 2.791l3.3-.466l.704-3.258zm-8.242 5.857l-2.596-2.791l-3.3.467l-.703 3.257l2.596 2.791l3.299-.466l.704-3.258zm-2.089-11.86l-2.596-2.791l-3.3.466l-.704 3.258l2.596 2.791l3.3-.467l.704-3.257z"/>`),
		g.Group(children),
	)
}