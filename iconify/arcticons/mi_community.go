package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiCommunity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.446 27.163L9.182 32.3c-.812 3.033 2.147 4.406 3.588 1.908c0 0 1.625 6.456 11.134 6.456s11.66-6.58 11.66-6.58c.643 2.4 3.653 1.335 3.653-.243c0-.59-1.377-6.631-1.377-6.631c1.224-.328 1.338-.986.463-1.862c.392-3.42-.176-5.998-2.296-9.293c2.193-1.267 6.493-3.48 6.493-6.162c0-2.44-2.355-3.147-4.218-2.07c-2.308 1.627-3.904 3.727-5.25 6.06c0 0-3.364-2.455-9.01-2.455s-8.706 2.335-8.706 2.335c-1.665-2.414-5.021-7.898-8.798-5.9c-2.94 2.14.942 6.106 5.828 8.064c-2.346 3.226-2.38 5.413-2.38 9.5c-.702.704-.633 1.436.48 1.735h0Z"/>`),
		g.Group(children),
	)
}