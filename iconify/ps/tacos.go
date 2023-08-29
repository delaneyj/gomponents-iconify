package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tacos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M299 149h21v22h-21v-22zm-107 0h21v22h-21v-22zm85 43h22v21h-22v-21zm64 0h22v21h-22v-21zm-192 0h22v21h-22v-21zm64 0h22v21h-22v-21zm288 4q7-24-6-49q-13-21-38-30l-30-6q-3-25-24-43q-25-21-55-10l-32 6l2 6q-32-6-62-6q-36 0-60 9v-9l-32-6q-26-7-55 10q-21 18-24 43l-27 6q-26 7-40 29.5T11 194l6 30Q0 270 0 320q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30q0-53-19-94zM43 320q0-51 21-90q26-55 77.5-89T256 107q62 0 114 33.5t78 87.5q21 45 21 90H43v2z"/>`),
		g.Group(children),
	)
}