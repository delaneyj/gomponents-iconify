package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M528 67q0-28-18.5-46T464 3H80Q53 3 34.5 21T16 67v245l-9 6l9 11v36q0 28 18.5 46T80 429h384q27 0 45.5-18t18.5-46v-98l9-10l-9-9V67zm-43 298q0 22-21 22H80q-21 0-21-22v-32l128-91l134 89l24-36l-47-32l83-83l104 87v98zm0-151l-108-90l-118 118l-72-52l-128 92V67q0-22 21-22h384q21 0 21 22v147zM208 109q0 18-12.5 30.5T165 152q-17 0-29.5-12.5T123 109q0-17 12.5-29.5T165 67q18 0 30.5 12.5T208 109z"/>`),
		g.Group(children),
	)
}