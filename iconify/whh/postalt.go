package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855.048 768h-87l-256 256V768h-343q-57 0-113-57t-56-115V172q0-58 56-115t113-57h686q57 0 113 57t56 115v424q0 58-56 115t-113 57zm-87-576h-512q-27 0-45.5 19t-18.5 45.5t18.5 45t45.5 18.5h512q26 0 45-18.5t19-45t-19-45.5t-45-19zm0 256h-512q-27 0-45.5 19t-18.5 45.5t18.5 45t45.5 18.5h512q26 0 45-18.5t19-45t-19-45.5t-45-19z"/>`),
		g.Group(children),
	)
}