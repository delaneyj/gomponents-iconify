package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feEject0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEject1" fill="currentColor"><path id="feEject2" d="M7 16h10a1 1 0 0 1 0 2H7a1 1 0 0 1 0-2Zm5.973-9.565l4.83 6.048c.359.448.214 1.054-.324 1.353a1.34 1.34 0 0 1-.648.164H7.169C6.523 14 6 13.563 6 13.024c0-.193.068-.381.196-.541l4.831-6.048c.358-.449 1.084-.57 1.622-.271c.128.071.238.163.324.27Z"/></g></g>`),
		g.Group(children),
	)
}