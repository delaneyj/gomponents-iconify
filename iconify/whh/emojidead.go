package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojidead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM374 377q9-10 9-24t-9.5-24t-23.5-10t-24 10l-39 39l-38-39q-10-10-24-10t-24 10t-10 24t10 24l39 38l-39 39q-10 10-10 24t10 23.5t24 9.5t24-9l38-39l39 39q10 9 24 9t23.5-9.5T383 478t-9-24l-39-39zm138 327q-53 0-90.5 28T384 800t37.5 68t90.5 28t90.5-28t37.5-68t-37.5-68t-90.5-28zm310-327q9-10 9-24t-9.5-24t-23.5-10t-24 10l-39 39l-38-39q-10-10-24-10t-24 10t-10 24t10 24l39 38l-39 39q-10 10-10 24t10 23.5t24 9.5t24-9l38-39l39 39q10 9 24 9t23.5-9.5T831 478t-9-24l-39-39z"/>`),
		g.Group(children),
	)
}