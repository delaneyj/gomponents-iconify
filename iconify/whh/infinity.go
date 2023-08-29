package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 568q-97 0-256-153q-159 153-256 153q-106 0-181-75T0 312t75-181t181-75q98 0 256 153Q671 56 768 56q106 0 181 75t75 181t-75 181t-181 75zM256 184q-57 0-92.5 37T128 312t35.5 91t92.5 37q22 0 46.5-13t44-32t35.5-38t25-32l9-13q-3-5-9-13.5t-24.5-31t-36.5-39t-43-30.5t-47-14zm512 0q-22 0-46.5 13t-44 32t-35.5 38t-25 32l-9 13q3 5 9 14t24.5 31t36.5 38.5t43 30.5t47 14q57 0 92.5-37t35.5-91t-35.5-91t-92.5-37z"/>`),
		g.Group(children),
	)
}