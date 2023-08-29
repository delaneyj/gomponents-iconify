package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M853.428 768h171l-352 256l-160-256h149l64-192h-341l192-576h192l-149 448h341zm-565-256q-13 0-22.5-9.5t-9.5-22.5V336q0-7-5-11.5t-11-4.5h-160q-7 0-11.5 4.5t-4.5 11.5v144q0 13-9.5 22.5t-22.5 9.5t-22.5-9.5t-9.5-22.5V128q0-53 37.5-90.5t90.5-37.5h64q53 0 90.5 37.5t37.5 90.5v352q0 13-9.5 22.5t-22.5 9.5zm-32-384q0-27-19-45.5t-45-18.5h-64q-27 0-45.5 18.5t-18.5 45.5v112q0 7 4.5 11.5t11.5 4.5h160q6 0 11-4.5t5-11.5V128z"/>`),
		g.Group(children),
	)
}