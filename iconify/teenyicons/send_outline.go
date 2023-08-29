package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m14.5.5l.46.197a.5.5 0 0 0-.657-.657L14.5.5Zm-14 6l-.197-.46a.5.5 0 0 0-.06.889L.5 6.5Zm8 8l-.429.257a.5.5 0 0 0 .889-.06L8.5 14.5ZM14.303.04l-14 6l.394.92l14-6l-.394-.92ZM.243 6.93l5 3l.514-.858l-5-3l-.514.858ZM5.07 9.757l3 5l.858-.514l-3-5l-.858.514Zm3.889 4.94l6-14l-.92-.394l-6 14l.92.394ZM14.146.147l-9 9l.708.707l9-9l-.708-.708Z"/>`),
		g.Group(children),
	)
}