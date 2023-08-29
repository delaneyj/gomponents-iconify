package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 327.618V164.133L165.067 0v161.756L0 327.618zM178.975 166.99v161.58l160.487-161.258l-160.487-.322zm-13.908 166.099V180.382L13.093 333.09h151.974zM0 348.249V511.26l163.126-163.01H0zm178.975-1.052V512l170.936-171.002v-165.56l-170.936 171.76z"/>`),
		g.Group(children),
	)
}