package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 806v122q0 40-28 68t-68 28t-68-28t-28-68q0 40-28 68t-68 28t-68-28t-28-68q0 40-28 68t-68 28t-68-28t-28-68V806q-89-63-140.5-156.5T0 448q0-91 38-174t102.5-143t153-95.5T480 0t186.5 35.5t153 95.5T922 274t38 174q0 108-51.5 202T768 806zM288 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm96 384h192l-96-128zm288-384q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}