package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dominoeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.56 960h-704q-53 0-90.5-37.5T.56 832V128q0-53 37.5-90.5T128.56 0h704q53 0 90.5 37.5t37.5 90.5v704q0 53-37.5 90.5t-90.5 37.5zm-672-896q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm320-640q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 640q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm320-640q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm0 320q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}