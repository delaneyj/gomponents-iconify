package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shirtbuttonalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM352 256q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm320-320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}