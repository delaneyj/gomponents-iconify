package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commentout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855.048 768h-215l-256 256V768h-215q-57 0-113-57t-56-115V172q0-58 56-115t113-57h686q57 0 113 57t56 115v424q0 58-56 115t-113 57zm-474-611q6-12 .5-20.5t-19.5-8.5h-68q-14 0-28 8.5t-20 20.5l-115 230q-3 6-3 13t3 13l115 198q6 12 20 20.5t28 8.5h68q14 0 19.5-8.5t-.5-20.5l-117-212zm195 3q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v256q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5V160zm0 384q0-13-9.5-22.5t-22.5-9.5h-64q-13 0-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5h64q13 0 22.5-9.5t9.5-22.5v-64zm320 0q0-13-9.5-22.5t-22.5-9.5h-192q-13 0-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5v-64z"/>`),
		g.Group(children),
	)
}