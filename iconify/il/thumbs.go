package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thumbs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M139 228q10 0 17 6t6 17v278q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V251q0-10 7-17t16-6h116zm485 0q22 0 41 9t32 26t18 36t0 43l-48 202q-8 34-35 54t-63 17l-180-15q-18-1-36-6l-111-32q-15-5-24-17t-10-28V248q0-20 16-34L443 14q21-18 46-9l4 1q17 6 26 21t5 33l-31 139q-2 11 5 20t18 9h108z"/>`),
		g.Group(children),
	)
}