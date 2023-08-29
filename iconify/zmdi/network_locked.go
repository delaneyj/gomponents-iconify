package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M395 197q-40 0-68 28.5T299 293v6q-22 19-22 48v64H0L405 5v193q-9-1-10-1zm53 128q9 0 15 6.5t6 15.5v85q0 9-6 15t-15 6H341q-8 0-14.5-6t-6.5-15v-85q0-9 6.5-15.5T341 325v-32q0-22 16-37.5t38-15.5t37.5 15.5T448 293v32zm-21 0v-32q0-13-9.5-22.5t-23-9.5t-22.5 9.5t-9 22.5v32h64z"/>`),
		g.Group(children),
	)
}