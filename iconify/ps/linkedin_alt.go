package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M434 5H41q-14 0-24 9.5T7 38v394q0 14 10 24t24 10h393q14 0 24-10t10-24V38q0-14-10-23.5T434 5zM147 390H77V182h70v208zm-35-236q-18 0-28.5-10.5T73 118t11-25.5T112 82q18 0 28.5 10.5T151 118t-10.5 25.5T112 154zm286 236h-70V279q0-47-35-47q-26 0-36 25q-2 5-2 17v116h-70q1-188 0-208h70v30q23-35 63-35q36 0 58 24t22 70v119zM255 213v-1v1z"/>`),
		g.Group(children),
	)
}