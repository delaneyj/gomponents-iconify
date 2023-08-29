package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roublealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm96-896H384q-26 0-45 19t-19 45v256h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64v64h-64q-26 0-45 19t-19 45.5t19 45t45 18.5h64v64q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5v-64h128q26 0 45-18.5t19-45t-19-45.5t-45-19H448v-64h160q93 0 158.5-65.5T832 352t-65.5-158.5T608 128zm0 320H448V256h160q40 0 68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}