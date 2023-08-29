package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M10 159q12 14 27 3q31-24 54-26q9 26 42 38v7q-37 8-61 37t-24 67v188q0 19 14 33t33 14h119q20 0 33.5-14t13.5-33V285q0-38-24-67t-61-37v-7q13-3 23-15l118 58q2 4 8 4q9 0 11-2q11-7 11-19V29q0-12-11-19t-21 0L197 68q-15-17-42-17q-21 0-37.5 12.5T95 95q-49 5-83 30q-8 7-8.5 17t6.5 17zm209-53l85-43v103l-85-43v-17zm0 367q0 4-5 4H95q-4 0-4-4V307h128v166zm-5-209H95q6-19 22.5-31t37.5-12q20 0 36.5 12t22.5 31zM155 93q9 0 15 6t6 16q0 9-6 15t-15 6q-10 0-16-6t-6-15q0-22 22-22z"/>`),
		g.Group(children),
	)
}